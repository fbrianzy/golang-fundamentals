# Golang Fundamentals

A hands-on, example-heavy introduction to Go (Golang). This repository is designed for self‑study **and** as a reference you can return to later. Each chapter is a folder with multiple runnable `.go` files plus a detailed `README.md` that explains concepts, shows annotated code, and proposes mini‑exercises.

> Target audience: motivated beginners to early‑intermediate developers who want a practical path to writing idiomatic Go.

## How this repo is organized

- Each numbered folder (e.g., `01-getting-started/`) focuses on a topic.
- Inside each folder:
  - Multiple `.go` files you can run directly with `go run`.
  - A `README.md` that explains the *why* and *how* behind the examples.
  - Optional subfolders (e.g., `internal/`, `cmd/`) for more realistic layouts.

## Quick start

```bash
# 1) Install Go (>= 1.21 recommended)
# 2) Clone this repo and enter the folder
git clone https://github.com/fbrianzy/golang-fundamentals.git
cd golang-fundamentals

# 3) Initialize a module name you prefer (edit go.mod if needed)
go mod tidy

# 4) Run any example
go run ./01-getting-started/hello.go

# 5) Run tests (chapters with tests)
go test ./10-testing-benchmarking/...
```

## Learning path

- [1. Getting Started](https://github.com/fbrianzy/golang-fundamentals/tree/main/01-getting-started)
- [2. Variables & Basic Types](https://github.com/fbrianzy/golang-fundamentals/tree/main/02-variables-and-types)
- [3. Control Flow (if, switch, loops)](https://github.com/fbrianzy/golang-fundamentals/tree/main/03-control-flow)
- [4. Functions & Multiple Returns](https://github.com/fbrianzy/golang-fundamentals/tree/main/04-functions) // We are here
- [5. Arrays, Slices, Maps]()
- [6. Structs, Methods, Interfaces]()
- [7. Packages & Modules]()
- [8. Error Handling & Panic/Recover]()
- [9. Concurrency: Goroutines & Channels]()
- [10. Testing & Benchmarking]()
- [11. I/O and Files]()
- [12. HTTP & REST (net/http)]()
- [13. Generics]()
- [14. Context & Cancellation]()
- [15. Reflection & Struct Tags]()
- [16. CLI Tool (Cobra‑style without deps)]()
- [17. Database (SQLite with `database/sql` + `modernc.org/sqlite`)]()
- [18. JSON & YAML]()
- [19. Logging & Configuration]()
- [20. Advanced Concurrency Patterns]()

> You don’t have to go in order—jump to what you need, then come back.

## Tips for success

- Keep `go fmt` and `go vet` in your muscle memory: `go fmt ./... && go vet ./...`
- Prefer small, composable functions. Leverage interfaces where it unlocks testing.
- Write tests early—even tiny ones—to learn how APIs behave.
- Read code from the standard library for idioms and patterns.

Happy hacking!
