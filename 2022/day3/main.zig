const std = @import("std");
const print = std.debug.print;

const UpperMax: usize = 52;
const LowerMax: usize = 26;

pub fn main(init: std.process.Init) !void {
    const io = init.io;

    const file = try std.Io.Dir.cwd().openFile(io, "input.txt", .{});
    defer file.close(io);

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &buf);
    var reader = &file_reader.interface;

    var total_priority1: usize = 0;
    var total_priority2: usize = 0;

    var line_count: u8 = 0;

    var line_bufs: [3][128]u8 = undefined;
    var line_lens: [3]usize = undefined;

    while (try reader.takeDelimiter('\n')) |raw_line| {
        const line = std.mem.trimEnd(u8, raw_line, "\r");

        // Part 1
        const k = line.len / 2;
        const left = line[0..k];
        const right = line[k..line.len];
        for (left) |lc| {
            if (std.mem.indexOfScalar(u8, right, lc)) |_| {
                total_priority1 += get_priority(lc);
                break;
            }
        }

        const idx = line_count % 3;
        @memcpy(line_bufs[idx][0..line.len], line);
        line_lens[idx] = line.len;
        line_count += 1;

        if (line_count == 3) {
            const l0 = line_bufs[0][0..line_lens[0]];
            const l1 = line_bufs[1][0..line_lens[1]];
            const l2 = line_bufs[2][0..line_lens[2]];

            for (l0) |c| {
                const contains_char = std.mem.indexOfScalar(u8, l1, c) != null and
                    std.mem.indexOfScalar(u8, l2, c) != null;
                if (contains_char) {
                    total_priority2 += get_priority(c);
                    break;
                }
            }
            line_count = 0;
        }
    }
    print("Part 1 -> Total Priority: {d}\n", .{total_priority1});
    print("Part 2 -> Total Priority: {d}\n", .{total_priority2});
}

fn get_priority(item: u8) usize {
    var priority: usize = 0;
    if ('A' <= item and item <= 'Z') {
        priority = UpperMax - ('Z' - item);
    } else {
        priority = LowerMax - ('z' - item);
    }
    return priority;
}
