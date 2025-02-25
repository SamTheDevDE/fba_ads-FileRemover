# FAFRemover

A Windows CLI tool to remove Facebook Ads files from your system. This tool requires administrator privileges to run.

![FAFRemover in action](assets/finished.png)

## Download

1. Go to the [Releases](https://github.com/yourusername/FAFRemover/releases) tab
2. Download the latest `FAFRemover.exe` file
3. Run the executable as administrator

## Building from Source

### Prerequisites

- Go 1.20 or higher
- Git

### Steps to Build

1. Clone the repository:
```bash
git clone https://github.com/yourusername/FAFRemover.git
cd FAFRemover
```

2. Install dependencies:
```bash
go get github.com/schollz/progressbar/v3
go get golang.org/x/sys/windows
```

3. Build the executable:
```bash
go build -o FAFRemover.exe
```

## Usage

Run the program as administrator with the following syntax:
```bash
FAFRemover.exe --dir=C:
```

Replace `C:` with the drive or directory you want to scan.

## Features

- Administrator privileges check
- Recursive directory scanning
- Progress bar with estimated time remaining
- Detailed completion summary
- Safe file deletion with error handling

## Requirements

- Windows operating system
- Administrator privileges

## License

MIT License