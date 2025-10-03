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
- [4. Functions & Multiple Returns](https://github.com/fbrianzy/golang-fundamentals/tree/main/04-functions)
- [5. Arrays, Slices, Maps](https://github.com/fbrianzy/golang-fundamentals/tree/main/05-collections)
- [6. Structs, Methods, Interfaces](https://github.com/fbrianzy/golang-fundamentals/tree/main/06-structs-methods-interfaces)
- [7. Packages & Modules](https://github.com/fbrianzy/golang-fundamentals/tree/main/07-packages-and-modules)
- [8. Error Handling & Panic/Recover](https://github.com/fbrianzy/golang-fundamentals/tree/main/08-errors-and-panic)
- [9. Concurrency: Goroutines & Channels](https://github.com/fbrianzy/golang-fundamentals/tree/main/09-concurrency-basics)
- [10. Testing & Benchmarking](https://github.com/fbrianzy/golang-fundamentals/tree/main/10-testing-and-benchmarking)
- [11. I/O and Files](https://github.com/fbrianzy/golang-fundamentals/tree/main/11-io-and-files)
- [12. HTTP & REST (net/http)](https://github.com/fbrianzy/golang-fundamentals/tree/main/12-http-and-rest)
- [13. Generics](https://github.com/fbrianzy/golang-fundamentals/tree/main/13-generics)
- [14. Context & Cancellation](https://github.com/fbrianzy/golang-fundamentals/tree/main/14-context-and-cancellation)
- [15. Reflection & Struct Tags](https://github.com/fbrianzy/golang-fundamentals/tree/main/15-reflection-and-tags)
- [16. CLI Tool (Cobra‑style without deps)](https://github.com/fbrianzy/golang-fundamentals/tree/main/16-cli-tool)
- [17. Database (SQLite with `database/sql` + `modernc.org/sqlite`)](https://github.com/fbrianzy/golang-fundamentals/tree/main/17-database-sqlite)
- [18. JSON & YAML](https://github.com/fbrianzy/golang-fundamentals/tree/main/18-json-and-yaml) // We are here
- [19. Logging & Configuration]()
- [20. Advanced Concurrency Patterns]()

> You don’t have to go in order—jump to what you need, then come back.

## Tips for success

- Keep `go fmt` and `go vet` in your muscle memory: `go fmt ./... && go vet ./...`
- Prefer small, composable functions. Leverage interfaces where it unlocks testing.
- Write tests early—even tiny ones—to learn how APIs behave.
- Read code from the standard library for idioms and patterns.

Happy hacking!
