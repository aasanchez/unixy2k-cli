# UnixY2K CLI

## Overview

UnixY2K is a command-line interface (CLI) tool that helps users track the time remaining until the infamous UnixY2K problem, a critical moment in computing history when 32-bit Unix timestamps will reach their maximum value on January 19, 2038, at 03:14:07 UTC.

## The UnixY2K Problem

The UnixY2K problem, often referred to as the "Year 2038 problem" or "Millennium bug," is a potential computer software issue where time representations could fail due to integer overflow in 32-bit systems. When the Unix timestamp reaches 2^31 - 1 (2,147,483,647), it will cause potential system failures, data corruption, and unexpected behavior in legacy systems.

## Features

- 🕒 Displays remaining time until the UnixY2K timestamp
- 📊 Multiple output formats
- 🔄 Continuous watch mode
- 📝 Detailed help information

## Prerequisites

- [Zig compiler](https://ziglang.org/download/) (version 0.11.0 or later)
- A Unix-like operating system (Linux, macOS, BSD)

## Installation

### Compiling from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/unixy2k.git
   cd unixy2k
   ```

2. Build the project:
   ```bash
   zig build-exe unixy2k.zig
   ```

3. (Optional) Install system-wide:
   ```bash
   sudo cp unixy2k /usr/local/bin/
   ```

## Usage

### Basic Usage

Run the tool without arguments to get a standard time remaining output:

```bash
./unixy2k
```

Example output:
```
Time remaining until UnixY2K: 14 years, 3 months, 22 days, 15 hours, 40 minutes, 15 seconds.
```

### Command-Line Options

| Option           | Short | Description                                      |
|-----------------|-------|--------------------------------------------------|
| `--help`        | `-h`  | Display help message and usage instructions      |
| `--simple`      | `-s`  | Display time in compact YY:MM:DD-HH:mm:ss format |
| `--watch`       | `-w`  | Continuously update and display remaining time   |

### Examples

1. Display help:
   ```bash
   ./unixy2k --help
   ```

2. Simple format:
   ```bash
   ./unixy2k --simple
   # Output: 14:03:22-15:40:15
   ```

3. Watch mode:
   ```bash
   ./unixy2k --watch
   # Continuously updates remaining time
   ```

## Technical Details

- **Language**: Zig
- **Timestamp Handling**: Uses 64-bit timestamp to calculate remaining time
- **Precision**: Calculates years, months, days, hours, minutes, and seconds
- **Constant**: `UNIXY2K_TIMESTAMP` set to 2,147,483,647 (2^31 - 1)

## Potential Limitations

- Assumes Gregorian calendar calculations
- Does not account for leap seconds
- Simplified time calculations might have minor inaccuracies

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Building and Testing

### Build

```bash
zig build
```

### Run Tests

```bash
zig test unixy2k.zig
```

## Related Resources

- [Unix Timestamp Converter](https://www.unixtimestamp.com/)
- [Year 2038 Problem - Wikipedia](https://en.wikipedia.org/wiki/Year_2038_problem)

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [Your Email]

Project Link: [https://github.com/yourusername/unixy2k](https://github.com/yourusername/unixy2k)

## Acknowledgments

- [Zig Programming Language](https://ziglang.org/)
- All contributors who help raise awareness about the UnixY2K problem

---

**Disclaimer**: This tool is for educational and awareness purposes. Always consult professional system administrators for critical system updates.
