const std = @import("std");
const root = @import("src/root.zig");

pub fn build(b: *std.Build) void {
    const exe = root.build(b);
    exe.install();
}
