# ComplianceGuard

A cross-platform CLI application built with Go, supporting macOS, Linux, and Windows on both AMD64 and ARM64 architectures.

## Features

- ✅ Cross-platform support (macOS, Linux, Windows)
- ✅ Multi-architecture (AMD64, ARM64)
- ✅ Simple and extensible command structure
- ✅ Built with Go standard library (no external dependencies)
- ✅ Easy-to-use build system

## Installation

### macOS (Homebrew)

```bash
# Add the tap (replace YOUR_USERNAME with your GitHub username)
brew tap YOUR_USERNAME/tap

# Install ComplianceGuard
brew install complianceguard
```

See [HOMEBREW_PUBLISHING.md](HOMEBREW_PUBLISHING.md) for details on publishing to Homebrew.

### From Binary Releases

Download pre-built binaries from the [Releases](https://github.com/YOUR_USERNAME/ComplianceGuard/releases) page.

### Build from Source

#### Prerequisites

- Go 1.25.7 or higher
- Make (optional, for using Makefile)

#### Building

#### Option 1: Using Make

```bash
# Build for your current platform
make build

# Build for all supported platforms
make build-all

# Build for a specific platform
make build-linux-amd64
make build-darwin-arm64
make build-windows-amd64
```

#### Option 2: Using build script

```bash
# Make the script executable
chmod +x build.sh

# Build for all platforms
./build.sh all

# Build for current platform
./build.sh current

# Build for specific platform
./build.sh linux amd64
./build.sh darwin arm64
./build.sh windows amd64
```

#### Option 3: Using Go directly

```bash
# Build for current platform
go build -o build/complianceguard .

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o build/complianceguard-linux-amd64 .
```

## Usage

```bash
# Show help
./complianceguard help

# Show version
./complianceguard version

# Display system information
./complianceguard info

# Run compliance checks
./complianceguard check

# Run with verbose output
./complianceguard check --verbose
./complianceguard check -v
```

## Supported Platforms

The application can be built for the following platforms:

| OS      | Architecture | Binary Name                        |
|---------|--------------|-------------------------------------|
| macOS   | AMD64        | complianceguard-darwin-amd64       |
| macOS   | ARM64        | complianceguard-darwin-arm64       |
| Linux   | AMD64        | complianceguard-linux-amd64        |
| Linux   | ARM64        | complianceguard-linux-arm64        |
| Windows | AMD64        | complianceguard-windows-amd64.exe  |
| Windows | ARM64        | complianceguard-windows-arm64.exe  |

## Development

### Running the Application

```bash
# Using Make
make run ARGS="version"
make run ARGS="check --verbose"

# Using Go directly
go run . version
go run . check --verbose
```

### Running Tests

```bash
make test
# or
go test -v ./...
```

### Project Structure

```
ComplianceGuard/
├── main.go              # Application entry point
├── cmd/
│   └── root.go         # CLI commands implementation
├── go.mod              # Go module definition
├── Makefile            # Build automation
├── build.sh            # Cross-platform build script
├── .gitignore          # Git ignore rules
└── README.md           # This file
```

### Adding New Commands

To add a new command, edit `cmd/root.go` and add a new case in the `Execute()` function:

```go
case "mycommand":
    return myCommandFunction()
```

Then implement your command function:

```go
func myCommandFunction() error {
    fmt.Println("Executing my command...")
    // Your logic here
    return nil
}
```

## Building for Distribution

To create binaries for distribution:

```bash
# Build all platforms
make build-all

# Binaries will be in the build/ directory
ls -lh build/
```

You can then distribute the appropriate binary for each platform to your users.

## Environment Variables

Currently, the application doesn't use environment variables, but you can easily add support by using `os.Getenv()` in your command implementations.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Troubleshooting

### Build Issues

If you encounter build issues:

1. Ensure you have Go 1.21 or higher: `go version`
2. Update dependencies: `make deps` or `go mod tidy`
3. Clean build artifacts: `make clean`

### Runtime Issues

For runtime issues:
- Use the `--verbose` flag to get more detailed output
- Check that you're using the correct binary for your platform

## Next Steps

Consider extending this CLI with:
- [ ] Configuration file support (YAML/JSON)
- [ ] Logging to files
- [ ] More sophisticated compliance checks
- [ ] Integration with external APIs
- [ ] Unit and integration tests
- [ ] GitHub Actions for automated builds
- [ ] Homebrew formula for macOS installation
- [ ] Debian/RPM packages for Linux distribution
