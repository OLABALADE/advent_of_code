const std = @import("std");
const print = std.debug.print;

const Range = struct {
    start: u32,
    end: u32,

    pub fn contains(self: Range, other: Range) bool {
        return self.start <= other.start and self.end >= other.end;
    }

    pub fn overlap(self: Range, other: Range) bool {
        if (self.contains(other)) return true;
        return self.start <= other.start and other.start <= self.end;
    }
};

pub fn main(init: std.process.Init) !void {
    const io = init.io;

    const file = try std.Io.Dir.cwd().openFile(io, "input.txt", .{});
    defer file.close(io);

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(io, &buf);
    var reader = &file_reader.interface;

    var count: usize = 0;
    var overlap_count: usize = 0;

    while (try reader.takeDelimiter('\n')) |raw_line| {
        const line = std.mem.trimEnd(u8, raw_line, "\r");
        const left, const right = try parseLine(line);

        if (left.contains(right) or right.contains(left)) {
            count += 1;
        }

        if (left.overlap(right) or right.overlap(left)) {
            overlap_count += 1;
        }
    }
    print("Part 1 -> Total pairs: {d}\n", .{count});
    print("Part 2 -> Total overlaping pairs: {d}\n", .{overlap_count});
}

fn parseLine(line: []const u8) ![2]Range {
    var tok = std.mem.tokenizeScalar(u8, line, ',');

    const first_tok = tok.next() orelse return error.MissingFirstRange;
    const second_tok = tok.next() orelse return error.MissingSecondRange;

    return .{
        try parseRange(first_tok),
        try parseRange(second_tok),
    };
}

fn parseRange(token: []const u8) !Range {
    const idx = std.mem.indexOfScalar(u8, token, '-') orelse return error.InvalidFormat;
    const start = try std.fmt.parseInt(u32, token[0..idx], 10);
    const end = try std.fmt.parseInt(u32, token[idx + 1 ..], 10);
    if (start > end) return error.InvalidRange;

    return .{ .start = start, .end = end };
}
