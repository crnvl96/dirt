# Dirt

A CLI tool to check if you have uncommitted changes or unpushed commits in any of your Git repositories.

Dirt scans specified directories (and their subdirectories up to 2 levels deep) for Git repositories and reports any that have uncommitted changes or unpushed commits.

## Installation

### Using Go

```bash
go install github.com/crnvl96/dirt@latest
```

### From Releases

Download the latest binary from the [releases page](https://github.com/crnvl96/dirt/releases).

## Usage

```bash
dirt [flags]
```

### Flags

- `-t, --target string`: Target directories to scan (scans recursively up to 2 levels). If not specified, scans the current directory.

### Examples

Scan the current directory:

```bash
dirt
```

Scan specific directories:

```bash
dirt -t ~/config -t ~/Developer -t ~/.nvim
```

### Output

- **Clean repositories**: Displays "All repositories are clean!" in green.
- **Dirty repositories**: Lists repositories with uncommitted changes or unpushed commits in red, with details on what's dirty.

## Building from Source

```bash
git clone https://github.com/crnvl96/dirt.git
cd dirt
go build -o dirt .
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
