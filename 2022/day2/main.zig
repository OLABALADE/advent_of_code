const std = @import("std");
const print = std.debug.print;

pub fn main(init: std.process.Init) !void {
    const io = init.io;
    const allocator = init.gpa;

    const file = try std.Io.Dir.cwd().openFile(io, "input.txt", .{});
    defer file.close(io);

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &buf);
    var reader = &file_reader.interface;

    const winScore = struct { wins: []const u8, hand_score: usize, equal_hand: []const u8 };

    var win_map: std.StringHashMap(winScore) = .init(allocator);
    defer win_map.deinit();

    const Hand = struct {
        draw: []const u8,
        win: []const u8,
        lose: []const u8,
        score: u8,
    };
    var hand_map: std.StringHashMap(Hand) = .init(allocator);
    defer hand_map.deinit();

    try hand_map.put("A", .{ .win = "C", .draw = "A", .lose = "B", .score = 1 }); // Rock
    try hand_map.put("B", .{ .win = "A", .draw = "B", .lose = "C", .score = 2 }); // Paper
    try hand_map.put("C", .{ .win = "B", .draw = "C", .lose = "A", .score = 3 }); // Scissors

    var alt_hand: std.StringHashMap([]const u8) = .init(allocator);
    defer alt_hand.deinit();

    try alt_hand.put("X", "A");
    try alt_hand.put("Y", "B");
    try alt_hand.put("Z", "C");

    var round: [2][]const u8 = undefined;
    var part1_score: usize = 0;
    var part2_score: usize = 0;

    while (try reader.takeDelimiter('\n')) |raw_line| {
        const line = std.mem.trimEnd(u8, raw_line, "\r");

        if (std.mem.cutScalar(u8, line, ' ')) |parts| {
            round[0] = parts[0];
            round[1] = parts[1];
        }

        const hand = hand_map.get(round[0]).?;
        const player_hand = alt_hand.get(round[1]).?;

        // Part 1
        if (std.mem.eql(u8, hand.lose, player_hand)) { //Win
            const win_score = hand_map.get(hand.lose).?.score;
            part1_score += (win_score + 6);
        } else if (std.mem.eql(u8, hand.draw, player_hand)) { // Draw
            part1_score += (hand.score + 3);
        } else {
            const lose_score = hand_map.get(hand.win).?.score; //Lose
            part1_score += lose_score;
        }

        // Part 2
        if (std.mem.eql(u8, round[1], "X")) { //Lose
            const losing_hand = hand_map.get(hand.win).?;
            part2_score += (losing_hand.score);
        } else if (std.mem.eql(u8, round[1], "Y")) { // Draw
            part2_score += (hand.score + 3);
        } else {
            const wining_hand = hand_map.get(hand.lose).?; // Win
            part2_score += (wining_hand.score + 6);
        }
    }

    std.debug.print("Part1: {d}\n", .{part1_score});
    std.debug.print("Part2: {d}\n", .{part2_score});
}
