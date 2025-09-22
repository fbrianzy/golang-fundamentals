# 02. Variables & Basic Types

Covers declarations (`var`, short `:=`), integers/floats, strings, and booleans. Understand zero values and explicit vs inferred types.

## Highlights
- Short declarations are great in local scopes.
- Strings are byte slices under the hood; slicing indexes bytes, not runes.
- Use `rune` to iterate Unicode‑aware with `for _, r := range s`.

## Files
- `numbers.go` — int/float basics.
- `strings.go` — slicing and length caveats.
- `booleans.go` — logical operators.

## Try
- Format numbers with `fmt.Printf` verbs.
- Count *runes* in a string using `utf8.RuneCountInString`.
