const std = @import("std");

pub fn main(init: std.process.Init) !void {
    const io = init.io;

    var file = try std.Io.Dir.cwd().openFile(io, "input.txt", .{});
    defer file.close(io);

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &buf);
    var reader = &file_reader.interface;

    var best: usize = 0;
    var sum: usize = 0;

    var top = [3]usize{ 0, 0, 0 };

    while (try reader.takeDelimiter('\n')) |raw_line| {
        const line = std.mem.trimEnd(u8, raw_line, "\r");
        if (line.len != 0) {
            const num = try std.fmt.parseInt(usize, line, 10);
            sum += num;
        } else {
            if (sum > top[0]) {
                top[0] = sum;
                std.mem.sort(usize, &top, {}, std.sort.asc(usize));
            }
            best = @max(sum, best);
            sum = 0;
        }
    }

    if (sum > 0) {
        if (sum > top[0]) {
            top[0] = sum;
            std.mem.sort(usize, &top, {}, std.sort.asc(usize));
        }
        best = @max(sum, best);
        sum = 0;
    }

    std.debug.print("Part 1: Max calories: {d}\n", .{best});

    var topScore: usize = 0;
    for (top) |n| {
        topScore += n;
    }
    std.debug.print("Part 2: Max calories: {d}\n", .{topScore});
}
