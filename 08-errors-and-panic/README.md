# 08. Error Handling, Panic & Recover

Prefer explicit `error` returns. Use `%w` to wrap errors. Reserve `panic` for truly unrecoverable states.

## Files
- `wrap_errors.go` — wrapping and `errors.Is`.
- `panic_recover.go` — controlled recovery demo.
