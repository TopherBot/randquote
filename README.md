# randquote

[![CI](https://github.com/youruser/randquote/actions/workflows/ci.yml/badge.svg)](https://github.com/youruser/randquote/actions)

`randquote` is a one‑file Go program that prints a random quote from a bundled list.

## Features
- No external dependencies – just the Go standard library.
- Uses the `embed` package to ship the quotes with the binary.
- A tiny test suite (`go test ./...`).
- Ready‑to‑use GitHub Actions CI (see the badge above).

## Installation
```bash
# Requires Go 1.23 or newer
go install github.com/youruser/randquote@latest
```

## Usage
```bash
randquote            # prints a random quote
randquote -n 3       # prints three random quotes
```

## Adding Your Own Quotes
Edit the `quotes.txt` file located in the repository root, one quote per line. The binary will automatically embed the updated list on the next build.

## Development
```bash
# Run tests
go test ./...

# Build locally
go build -o randquote ./
```

---
*This project follows the best‑practice checklist for GitHub Actions security: actions are pinned to commit SHAs, the default `GITHUB_TOKEN` permissions are set to read‑only, and no secrets are used in the workflow.*
