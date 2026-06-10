# AGENTS.md

Guidelines and context for AI agents working on the codebase.

## LLM Guidelines

### Do
- Run `go build` after any code change to verify it compiles
- Run `go test ./...` before committing to ensure tests pass
- Keep existing functionality unchanged by default
- Test changes in browser when modifying UI (index.html)
- **Parallelize operations where possible** — use goroutines, worker pools, batch processing
- **Prefer caching** — cache computed results, API responses, and expensive operations; use TTL-based invalidation
- Commit changes locally after task completion

### Don't
- Remove existing features or API endpoints
- Change external API behavior
- Push commits to remote — commit locally only, never push

### Ask First
- Refactoring that changes behavior
- Adding new dependencies
- Removing functionality
- Breaking changes to data formats

## Tech Stack

- **Language:** Go 1.25 (pure Go, no CGO/MinGW required)


## Developer Tools

The following tools are installed in `~/go/bin/` and require `export PATH="$HOME/go/bin:$PATH"` to run:

| Tool | Purpose | Run |
|------|---------|-----|
| `air` | Live-reload development | `air` (uses `.air.toml`) |
| `golangci-lint` | Comprehensive linter | `golangci-lint run ./...` |
| `staticcheck` | Static analysis | `staticcheck ./...` |
| `gofumpt` | Stricter `gofmt` | `gofumpt -w .` |
| `gocyclo` | Cyclomatic complexity | `gocyclo -over 25 ./...` |
| `dupl` | Duplicate code detection | `dupl -t 100 ./...` |
| `dlv` | Delve debugger | `dlv debug ./cmd/thaifa` |
| `govulncheck` | Vulnerability scanner | `govulncheck ./...` |
| `gotests` | Test boilerplate generator | `gotests -all -w file.go` |
| `richgo` | Colorized test output | `richgo test ./...` |

## Browse Tool — Agent Reference

### Overview

The `scripts/browse` tool is a lightweight wrapper around Chromium that fetches webpage content while bypassing bot detection/WAF protection.

### When to Use

Use `browse` instead of `curl`/`wget` when:
- Site returns "Access Denied" or "Blocked"
- Site uses Cloudflare, Akamai, or similar WAF
- Site is a Single Page Application (SPA) that requires JavaScript rendering
- Standard HTTP tools return empty or bot challenge pages

### Usage

```bash
# Basic usage
browse "https://example.com"

# Save to file
browse "https://example.com" > output.html

# Pipe to other tools
browse "https://example.com" | grep "pattern"

# For JSON APIs
browse "https://api.example.com/data" | jq '.'
```

## Code Conventions

### Go Style
- Standard Go formatting (gofmt)
- No comments unless explicitly requested
- Error handling with fmt.Errorf and %w for wrapping
- Mutex-protected settings with RLock/Lock
