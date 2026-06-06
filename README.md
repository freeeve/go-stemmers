# go-stemmers

A Go port of the [`rust-stemmers`](https://github.com/CurrySoftware/rust-stemmers)
crate: the [Snowball](https://snowballstem.org/) stemming algorithms with a small
`Algorithm`/`Stemmer` API. Every algorithm produces output **identical** to
rust-stemmers 1.2.0 (and thus to the reference Snowball implementation), verified
word-for-word against the canonical Snowball vocabularies.

```go
import "github.com/freeeve/go-stemmers"

s := stemmers.New(stemmers.English)
s.Stem("fruitlessly") // "fruitless"
s.Stem("national")    // "nation"
```

Input is expected to be **already lowercased**, matching the Snowball/rust-stemmers
contract.

## Why output parity matters

This library exists so a Go service can stem text exactly as a Rust/wasm peer does.
In an inverted index, the indexed term and the query term must reduce to the *same*
stem byte-for-byte, or every stemmed lookup silently misses. go-stemmers is the
build-side mirror of rust-stemmers used by [roaringrange](https://github.com/freeeve/roaringrange)'s
term index: build the index from Go, query it from the wasm reader, and the
stemming agrees.

Note this is **output-string parity**, not a serialized byte format — a stemmer has
no on-disk representation. The contract is `Stem(w)` returns the identical string
rust-stemmers returns, for every `w`.

## Algorithms

All 18 rust-stemmers algorithms are ported:

`Arabic`, `Danish`, `Dutch`, `English`, `Finnish`, `French`, `German`, `Greek`,
`Hungarian`, `Italian`, `Norwegian`, `Portuguese`, `Romanian`, `Russian`,
`Spanish`, `Swedish`, `Tamil`, `Turkish`.

The `Algorithm` enum mirrors rust-stemmers' order one-to-one.

## How it's structured

- `internal/snowball` — the runtime (`Env`, `Among`, `FindAmong`/`FindAmongB`): a
  byte-faithful port of rust-stemmers' `snowball_env.rs`/`among.rs`. Shared by
  every algorithm.
- `internal/<lang>` — one package per language, each a faithful translation of the
  generated `algorithms/<lang>.rs` (Snowball labelled blocks → Go labelled loops,
  result-code dispatch → `switch`).
- `stemmers.go` — the public `Algorithm`/`Stemmer` API.
- `rustgen/` — a small Rust binary linking rust-stemmers 1.2.0; the **oracle** that
  generates goldens and backs the live differential test.
- `cmd/oacorpus/` — streams the public OpenAlex Works snapshot (CC0) into a word
  corpus for benchmarking; nothing is vendored.

## Testing

```sh
go test ./...                 # vector conformance + known cases + fuzz seed corpus
```

- **`TestVectors`** stems every word in each `testdata/voc_<lang>.txt` and asserts
  it equals `res_<lang>.txt` — the parity teeth. The goldens are rust-stemmers
  1.2.0 output (see [`testdata/NOTICE.md`](testdata/NOTICE.md) for provenance).
- **`FuzzStem`** checks no algorithm panics and that stems stay valid UTF-8 and
  never grow.
- **`TestRustgenParity`** is the cross-language differential check against the live
  oracle over a large OpenAlex corpus. It runs only when both `rustgen` is built
  and a corpus has been streamed (otherwise it skips, keeping CI hermetic):

  ```sh
  cargo build --release --manifest-path rustgen/Cargo.toml
  go run ./cmd/oacorpus -budget 200000 -out testdata/openalex.corpus
  go test -run TestRustgenParity
  ```

## Benchmarks

```sh
go test -bench=English -benchmem        # OpenAlex corpus if present, else the vocab
go test -bench=All                      # per-language throughput
```

## Updating the pinned rust-stemmers version

The parity contract is pinned to rust-stemmers 1.2.0 (`rustgen/Cargo.toml`).
Bumping it means regenerating the goldens for any language whose output changed;
the goldens rust-stemmers ships no vectors for (da, nl, hu, sv, ta, tr) are
produced by `rustgen` over snowball-data vocabularies — see `testdata/NOTICE.md`.

## License

MIT — see [LICENSE](LICENSE). Test fixtures under `testdata/` carry their own
provenance and licenses (documented in `testdata/NOTICE.md`); they are used only by
`go test` and are not part of the importable library.
