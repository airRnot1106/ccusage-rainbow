# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**ccusage-rainbow** is a Go CLI application that displays animated rainbow ASCII art text. It integrates with the `ccusage` tool to fetch and display Claude API usage costs as large, colorful terminal output. The project implements Clean Architecture with clear layer separation.

## Essential Commands

**Development Environment:**
```bash
nix develop                    # Enter development shell (provides Go, linters, formatters)
direnv allow                   # Auto-load development environment
```

**Build & Run:**
```bash
go run main.go                 # Run with real cost data from ccusage
go run main.go --test          # Run with test data ($9999.99)
go run main.go --custom --text "HELLO"  # Run with custom text
go build -o ccusage-rainbow    # Build binary
```

**Code Quality:**
```bash
golangci-lint run              # Run Go linter
gofumpt -w .                   # Format Go code
treefmt                        # Format all code (Go, Nix, etc.)
```

**Testing:**
```bash
go test ./...                  # Run all tests
go test -v ./internal/usecase/rainbow/  # Run specific package tests
```

## Architecture Overview

This project follows **Clean Architecture** with strict dependency inversion:

### Dependency Flow
```
CLI/TUI → Use Cases → Infrastructure
    ↓         ↓           ↓
Domain Interfaces ← Domain Entities
```

### Key Layers

**Domain (`/internal/domain/`):**
- `entities/` - Core business objects (Text, CostResponse, RainbowAnimation)
- `interfaces/` - Contracts for external dependencies (ASCIIRenderer, ColorAnimator, CostService)

**Use Cases (`/internal/usecase/`):**
- `rainbow/` - Core business logic for animated ASCII text rendering
- `cost/` - Cost data fetching and formatting logic

**Infrastructure (`/internal/infrastructure/`):**
- `ascii/` - ASCII art rendering with 3 font sizes (5x7, 7x9, 10x13)
- `color/` - Rainbow animation using 7-color cycling
- `cost/` - External integration with `npx ccusage@latest -j`

**Interface Adapters (`/internal/interfaces/`):**
- `cli/` - Cobra-based command-line interface
- `tui/` - Bubble Tea terminal user interface

**Frameworks (`/internal/frameworks/`):**
- `di/` - Manual dependency injection container

### Critical Implementation Details

**ASCII Rendering:**
- Uses Unicode block characters (`██`) which have display width of 2 but string length of 1
- MUST use `lipgloss.Width()` for accurate width calculations, never `len()`
- Three font sizes with dynamic selection based on terminal dimensions
- Special spacing logic for decimal points (tighter than normal character spacing)

**Animation System:**
- 7-color rainbow with character-offset cycling
- Uses Bubble Tea's tick-based animation with 100ms intervals
- Colors applied per-character using Lipgloss styling

**External Integration:**
- Executes `npx ccusage@latest -j` to fetch cost data
- Parses JSON response into domain entities
- Fallback to "ERROR" display if ccusage fails

## Development Guidelines

**Adding New Features:**
1. Start with domain entities and interfaces
2. Implement use cases with business logic
3. Create infrastructure implementations
4. Wire dependencies in DI container
5. Update CLI/TUI interfaces last

**Testing Approach:**
- Mock external dependencies using domain interfaces
- Test use cases independently of infrastructure
- Use dependency injection for testability

**Font/Character Modifications:**
- Character patterns must be consistent width within each font size
- Small: 7 chars wide, Medium: 9 chars wide, Large: 14 chars wide (with exceptions for special characters)
- Update all three font sizes when adding new characters
- Test with various terminal widths to ensure proper font selection

**Color/Animation Changes:**
- Rainbow uses exactly 7 colors in fixed order
- Character offset creates wave effect across text
- Animation timing controlled via Bubble Tea model updates