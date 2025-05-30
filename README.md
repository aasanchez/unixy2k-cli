# Unixy2K CLI (Go Edition)

## Overview

Unixy2K CLI is a Go-based terminal application that provides a real-time display of:
*   Current UTC (Coordinated Universal Time)
*   Current Epoch time (seconds since January 1, 1970, UTC)
*   Binary representation of the current Epoch time (32-bit)
*   A live countdown to the Year 2038 problem (Unix epoch rollover: 2038-01-19 03:14:07 UTC).

This tool utilizes the excellent Charm SH libraries (Bubble Tea and Lip Gloss) for its interactive and styled terminal user interface.

## Features

*   **Live Countdown:** Real-time updates for all displayed time values.
*   **Full-Screen Terminal UI:** An immersive, focused display.
*   **Comprehensive Time Info:** Shows UTC, Epoch, and the 32-bit binary representation of the Epoch.
*   **Cross-Platform:** Built with Go, easily compilable for various operating systems and architectures.

## Building and Running Locally

### Prerequisites

*   Go (Version 1.21 or newer recommended)

### Instructions

1.  **Clone the repository** (if you haven't already):
    ```bash
    # Replace <repo_url> with the actual repository URL
    # git clone <repo_url>
    # cd <repo_directory_name>
    ```
    (Note: For this environment, the code is likely already checked out in `/app`)

2.  **To run directly from source** (from the project root directory `/app`):
    ```bash
    go run github.com/aasanchez/unixy2k-cli
    ```
    This command fetches dependencies, compiles, and runs the application in one step. The main package is located in the `src/` directory.

3.  **To build the binary** (from the project root directory `/app`):
    This command will create an executable file named `unixy2k-cli` (or `unixy2k-cli.exe` on Windows) in the current directory.
    ```bash
    go build -o unixy2k-cli ./src
    ```
    Then run the compiled application:
    ```bash
    ./unixy2k-cli
    ```

## Cross-Compilation

Go excels at cross-compilation. You can build executables for different operating systems and architectures from your current machine by setting the `GOOS` (target Operating System) and `GOARCH` (target Architecture) environment variables.

**Target Operating Systems (`GOOS`):**
*   `linux` (for Ubuntu, Debian, AlmaLinux, OpenSuse, etc.)
*   `darwin` (for MacOS)
*   `windows`

**Target Architectures (`GOARCH`):**
*   `amd64` (Most common for x86-64 desktops and servers)
*   `arm64` (Apple Silicon Macs, newer Raspberry Pi models, some ARM-based servers)

**Example Commands (run from the project root directory):**

First, you might need to create a directory for your builds:
```bash
mkdir build
```

*   **For Linux:**
    *   AMD64:
        ```bash
        GOOS=linux GOARCH=amd64 go build -o build/unixy2k-cli-linux-amd64 ./src
        ```
    *   ARM64:
        ```bash
        GOOS=linux GOARCH=arm64 go build -o build/unixy2k-cli-linux-arm64 ./src
        ```

*   **For MacOS (Darwin):**
    *   AMD64:
        ```bash
        GOOS=darwin GOARCH=amd64 go build -o build/unixy2k-cli-darwin-amd64 ./src
        ```
    *   ARM64 (Apple Silicon):
        ```bash
        GOOS=darwin GOARCH=arm64 go build -o build/unixy2k-cli-darwin-arm64 ./src
        ```

*   **For Windows:**
    *   AMD64:
        ```bash
        GOOS=windows GOARCH=amd64 go build -o build/unixy2k-cli-windows-amd64.exe ./src
        ```

The compiled binaries will be placed in the `build/` directory.

## Usage

Run the application (either directly via `go run` or by executing a compiled binary).

*   Press `q`, `Ctrl+C`, or `Esc` to quit the application.

## License

This project is licensed under the MIT License.

## Acknowledgments

*   This application is built using the fantastic [Charm SH](https://charm.sh/) libraries:
    *   [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the terminal application framework.
    *   [Lip Gloss](https://github.com/charmbracelet/lipgloss) for terminal styling.
