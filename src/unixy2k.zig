const std = @import("std");
const time = std.time;

const UNIXY2K_TIMESTAMP: u64 = 2147483647; // Unix timestamp for 2038-01-19 03:14:07 UTC

// Function to display the help message
fn show_help() void {
    const bold = "\x1b[1m"; // Bold
    const reset = "\x1b[0m"; // Reset to default formatting

    const help_message = "{s}UnixY2K{s}\n\n" ++
        "A CLI tool to display the time remaining until the UnixY2K problem (January 19, 2038)\n\n" ++
        "{s}Usage:{s}\n" ++
        "unixy2k [--help | -h | --simple | -s | --watch | -w]\n\n" ++
        "{s}Options:{s}\n" ++
        "--help, -h   Display this help message\n" ++
        "--simple, -s Display the remaining time in the format YY:MM:DD-HH:mm:ss\n" ++
        "--watch, -w  Continuously display the remaining time, refreshing every second\n\n" ++
        "The UnixY2K problem will occur when the Unix timestamp reaches its maximum value (2^31 - 1)\n" ++
        "on January 19, 2038, at 03:14:07 UTC. This tool shows how much time remains until that moment.\n" ++
        "\n" ++
        "Visit https://unixy2k.com/\n\n";

    std.debug.print(help_message, .{ bold, reset, bold, reset, bold, reset });
}

// Function to calculate the time remaining
fn calculate_remaining_time() [6]u64 {
    const currentTimestamp = @as(u64, @intCast(time.timestamp()));
    const remainingSeconds = if (currentTimestamp < UNIXY2K_TIMESTAMP)
        UNIXY2K_TIMESTAMP - currentTimestamp
    else
        0;

    const secondsPerMinute = 60;
    const secondsPerHour = secondsPerMinute * 60;
    const secondsPerDay = secondsPerHour * 24;
    const secondsPerYear = secondsPerDay * 365;
    const secondsPerMonth = secondsPerYear / 12;

    const years = remainingSeconds / secondsPerYear;
    var remaining = remainingSeconds % secondsPerYear;

    const months = remaining / secondsPerMonth;
    remaining %= secondsPerMonth;

    const days = remaining / secondsPerDay;
    remaining %= secondsPerDay;

    const hours = remaining / secondsPerHour;
    remaining %= secondsPerHour;

    const minutes = remaining / secondsPerMinute;
    const seconds = remaining % secondsPerMinute;

    return [6]u64{ years, months, days, hours, minutes, seconds };
}

pub fn main() !void {
    var general_purpose_allocator = std.heap.GeneralPurposeAllocator(.{}){};
    const gpa = general_purpose_allocator.allocator();
    defer _ = general_purpose_allocator.deinit();

    var args = try std.process.ArgIterator.initWithAllocator(gpa);
    defer args.deinit();

    // Skip the first argument (program name)
    _ = args.skip();

    // Check if arguments are passed
    if (args.next()) |first_arg| {
        if (std.mem.eql(u8, first_arg, "--help") or std.mem.eql(u8, first_arg, "-h")) {
            show_help();
            return;
        } else if (std.mem.eql(u8, first_arg, "--simple") or std.mem.eql(u8, first_arg, "-s")) {
            const remaining_time = calculate_remaining_time();
            std.debug.print("{:0>2}:{:0>2}:{:0>2}-{:0>2}:{:0>2}:{:0>2}\n", .{
                remaining_time[0], // Years
                remaining_time[1], // Months
                remaining_time[2], // Days
                remaining_time[3], // Hours
                remaining_time[4], // Minutes
                remaining_time[5], // Seconds
            });
            return;
        } else if (std.mem.eql(u8, first_arg, "--watch") or std.mem.eql(u8, first_arg, "-w")) {
            while (true) {
                const remaining_time = calculate_remaining_time();
                std.debug.print("{:0>2}:{:0>2}:{:0>2}-{:0>2}:{:0>2}:{:0>2}\r", .{
                    remaining_time[0], // Years
                    remaining_time[1], // Months
                    remaining_time[2], // Days
                    remaining_time[3], // Hours
                    remaining_time[4], // Minutes
                    remaining_time[5], // Seconds,
                });
                std.io.getStdOut().writeAll("\r") catch {};
                std.time.sleep(1_000_000_000); // Sleep for 1 second
            }
        }
    }

    const remaining_time = calculate_remaining_time();
    std.debug.print("Time remaining until UnixY2K: {} years, {} months, {} days, {} hours, {} minutes, {} seconds.\n", .{
        remaining_time[0], // Years
        remaining_time[1], // Months
        remaining_time[2], // Days
        remaining_time[3], // Hours
        remaining_time[4], // Minutes
        remaining_time[5], // Seconds,
    });
}
