const std = @import("std");
const print = std.debug.print;

pub fn main(init: std.process.Init) !void {
    const io = init.io;
    const allocator = init.gpa;

    const max_bytes = 50 * 1024; // 50KB
    const raw_content = try std.Io.Dir.cwd().readFileAlloc(io, "input.txt", allocator, .limited(max_bytes));
    defer allocator.free(raw_content);

    var normalized: std.ArrayList(u8) = .empty;
    defer normalized.deinit(allocator);
    for (raw_content) |byte| {
        if (byte != '\r') {
            try normalized.append(allocator, byte);
        }
    }

    var tokens = std.mem.tokenizeSequence(u8, normalized.items, "\n\n");
    var warehouse: std.ArrayList([]const u8) = .empty;
    defer warehouse.deinit(allocator);

    while (tokens.next()) |tok| {
        try warehouse.append(allocator, tok);
    }

    if (warehouse.items.len < 2) {
        return error.InvalidInput;
    }

    var crate_section = warehouse.items[0];
    if (std.mem.lastIndexOfScalar(u8, crate_section, '\n')) |last_nl| {
        crate_section = crate_section[0..last_nl];
    }

    var crates1 = try parse_crates(allocator, crate_section);
    var crates2 = try parse_crates(allocator, crate_section);

    defer {
        for (crates1.items, crates2.items) |*inner1, *inner2| {
            inner1.deinit(allocator);
            inner2.deinit(allocator);
        }
        crates1.deinit(allocator);
        crates2.deinit(allocator);
    }

    const intructions = try parse_instructions(allocator, warehouse.items[1]);
    defer allocator.free(intructions);

    for (intructions) |instr| {
        try perform(allocator, instr, &crates1, false);
        try perform(allocator, instr, &crates2, true);
    }

    var message1: std.ArrayList(u8) = .empty;
    defer message1.deinit(allocator);

    var message2: std.ArrayList(u8) = .empty;
    defer message2.deinit(allocator);

    for (crates1.items, crates2.items) |crate1, crate2| {
        if (crate1.getLastOrNull()) |crate_name1| {
            try message1.append(allocator, crate_name1);
        }
        if (crate2.getLastOrNull()) |crate_name2| {
            try message2.append(allocator, crate_name2);
        }
    }

    print("Part1 -> Message: {s}\n", .{message1.items});
    print("Part2 -> Message: {s}\n", .{message2.items});
}

fn parse_crates(allocator: std.mem.Allocator, items: []const u8) !std.ArrayList(std.ArrayList(u8)) {
    var result: std.ArrayList(std.ArrayList(u8)) = .empty;

    errdefer {
        for (result.items) |*inner| {
            inner.deinit(allocator);
        }
        result.deinit(allocator);
    }

    var max_len: usize = 0;
    var size_iter = std.mem.tokenizeScalar(u8, items, '\n');
    while (size_iter.next()) |raw_tok| {
        const tok = std.mem.trimEnd(u8, raw_tok, "\r");
        if (tok.len > max_len) {
            max_len = tok.len;
        }
    }

    const num_crates = (max_len + 1) / 4;
    for (0..num_crates) |_| {
        const col: std.ArrayList(u8) = .empty;
        try result.append(allocator, col);
    }

    var iter = std.mem.tokenizeScalar(u8, items, '\n');
    while (iter.next()) |raw_tok| {
        const tok = std.mem.trimEnd(u8, raw_tok, "\r");
        var col_idx: usize = 0;

        while (col_idx < num_crates) : (col_idx += 1) {
            const start = col_idx * 4;
            if (start + 1 >= tok.len) break;
            if (tok[start] == '[') {
                try result.items[col_idx].append(allocator, tok[start + 1]);
            }
        }
    }

    for (result.items) |col| {
        std.mem.reverse(u8, col.items);
    }

    return result;
}

fn parse_instructions(allocator: std.mem.Allocator, items: []const u8) ![][3]u32 {
    var tokens = std.mem.tokenizeScalar(u8, items, '\n');
    var lines: std.ArrayList([3]u32) = .empty;

    while (tokens.next()) |tok| {
        const instr = try parse_inst_line(tok);
        try lines.append(allocator, instr);
    }
    return lines.toOwnedSlice(allocator);
}

fn parse_inst_line(line: []const u8) ![3]u32 {
    var iter = std.mem.tokenizeScalar(u8, line, ' ');
    var result: [3]u32 = .{ 0, 0, 0 };

    var i: usize = 0;
    while (iter.next()) |tok| {
        if (std.fmt.parseInt(u32, tok, 10)) |num| {
            result[i] = num;
            i += 1;
        } else |_| {
            continue;
        }
    }
    return result;
}

fn perform(allocator: std.mem.Allocator, instruction: [3]u32, crates: *std.ArrayList(std.ArrayList(u8)), isPart2: bool) !void {
    const num_crates, const source, const destination = instruction;
    const count: usize = @intCast(num_crates);

    if (source == 0 or destination == 0 or source > crates.items.len or destination > crates.items.len) {
        return error.InvalidStackIndex;
    }

    var src = &crates.items[source - 1];
    var dst = &crates.items[destination - 1];

    if (count > src.items.len) {
        return error.NotEnoughCrates;
    }

    const start = src.items.len - count;
    const moving = src.items[start..];

    if (isPart2) {
        try dst.appendSlice(allocator, moving);
    } else {
        var idx = moving.len;
        while (idx > 0) {
            idx -= 1;
            try dst.append(allocator, moving[idx]);
        }
    }

    src.items.len = start;
}
