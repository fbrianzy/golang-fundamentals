# 01. Getting Started

In this chapter you will install Go, verify your environment, and run your first programs.

## Key ideas
- `go version`, `go env GOPATH` to verify installation.
- Run a single file: `go run hello.go`.
- Build a binary: `go build -o app hello.go`.
- Pass command-line arguments via `os.Args`.
- Read environment variables via `os.Getenv`.

## Files
- `hello.go` — classic hello.
- `args.go` — prints command-line arguments.
- `env.go` — prints a selected environment variable.

## Exercises
1. Modify `hello.go` to greet a name from `os.Args`.
2. Add a flag (e.g., `-shout`) that uppercases the greeting (hint: `flag` package).
