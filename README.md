# ⚠️ Archived

This repository is archived and no longer actively maintained.  
Please see the new home of the project:

👉 https://github.com/airRnot1106/ccusage-gorgeous

# ccusage-rainbow 🌈

Tool to display ccusage in a gorgeous way

https://github.com/user-attachments/assets/7f3e5983-3f73-4515-b32a-18853f7cc896

## 🚀 Quick Start

### Using Nix (Recommended)

```bash
# Clone the repository
git clone https://github.com/your-username/ccusage-rainbow.git
cd ccusage-rainbow

# Enter development environment (sets up git hooks automatically)
nix develop

# Build binary
nix build

# Run
nix run
```

### Using Go

```bash
# Install dependencies
go mod download

# Run with your actual ccusage data
go run main.go

# Build binary
go build
```

## 🔄 Dependency Management

This project uses [Dependabot](https://docs.github.com/code-security/dependabot) for automated dependency updates:

- **Go modules**: Updated weekly on Mondays
- **Minor/patch updates**: Auto-merged after CI passes
- **Major updates**: Require manual review
- **Charm libraries**: Grouped separately for UI framework updates

## 🪝 Git Hooks

Pre-commit hooks are automatically installed when you run `nix develop`:

- **treefmt**: Automatic code formatting
- **golangci-lint**: Code linting
- **go test**: Run tests

This ensures code quality and consistency before commits are made.

## 📝 LICENSE

MIT

## 🙏 Acknowledgments

- [ccusage](https://github.com/ryoppippi/ccusage) - A CLI tool for analyzing Claude Code usage from local JSONL files.
