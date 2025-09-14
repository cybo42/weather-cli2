# Weather CLI

A simple command-line utility that prints "hello world".

## Installation

### Using Homebrew (macOS and Linux)

```bash
# Add the tap
brew tap cybo42/demo-homebrew-tap

# Install the CLI
brew install weather-cli
```

### Using DEB Package (Debian, Ubuntu, etc.)

```bash
# Download the latest .deb package from the releases page
curl -LO "https://github.com/cybo42/weather-cli/releases/latest/download/weather-cli_linux_amd64.deb"

# Install the package
sudo dpkg -i weather-cli_linux_amd64.deb
```

### Using RPM Package (Fedora, CentOS, RHEL, etc.)

```bash
# Download the latest .rpm package from the releases page
curl -LO "https://github.com/cybo42/weather-cli/releases/latest/download/weather-cli_linux_amd64.rpm"

# Install the package
sudo rpm -i weather-cli_linux_amd64.rpm
```

### From Source

```bash
# Clone the repository
git clone https://github.com/cybo42/weather-cli.git
cd weather-cli

# Build the binary
go build -o weather-cli

# Install the binary (optional)
go install
```

## Usage

Simply run the command:

```bash
weather-cli
```

This will print "hello world" to the console.

To check the version:

```bash
weather-cli --version
```

## Development

### Prerequisites

- Go 1.25 or higher
- Git

### Building

```bash
go build -o weather-cli
```

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) to automate the release process and [GitHub Actions](https://github.com/features/actions) to automate the workflow.

To create a new release:

1. Update the code as needed
2. Commit your changes
3. Tag the release:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   ```
4. Push the tag:
   ```bash
   git push origin v1.0.0
   ```

The GitHub Actions workflow will automatically:
1. Build the binaries for multiple platforms
2. Create a GitHub release with the binaries, RPM packages, and DEB packages
3. Update the Homebrew formula in the [demo-homebrew-tap](https://github.com/cybo42/demo-homebrew-tap)

### Testing a Release Locally

To test the release process locally without publishing:

```bash
goreleaser release --snapshot --clean
```

## License

MIT