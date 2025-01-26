const std = @import("std");
const main = @import("main");

pub fn build(b: *std.Build) void {
    const mode = b.standardReleaseOptions();
    const exe = b.addExecutable("unixy2k", "main.zig");
    exe.setBuildMode(mode);
    exe.install();
}

pub fn run() void {
    main.main();
}
