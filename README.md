## Mini Search Engine (Go)

Lightweight inverted-index search toy. It builds an in-memory index over small text docs and supports simple AND queries from stdin.

### How it works
- Reads `.txt` docs from `internal/data/docs`.
- Normalizes words (lowercase, strips punctuation) and builds an inverted index.
- Takes space-separated query terms from stdin.
- Returns document IDs that contain **all** queried terms and prints their contents.

### Getting started
1) Prereqs: Go 1.20+  
2) Install deps (none beyond stdlib).  
3) Run:
```bash
go run ./cmd/search
```
4) When prompted, type your search terms and hit Enter, e.g.:
```
data structures
```

### Project structure
- `cmd/search/main.go` — CLI entrypoint.
- `internal/index` — inverted index types + builder.
- `internal/search` — query evaluation (posting list intersection).
- `internal/helpers` — file IO and token sanitization.
- `internal/data/docs` — sample documents to index.

### Notes
- Queries are AND-only; terms not in the index yield no results.
- Update or add `.txt` files in `internal/data/docs` to change the corpus, then rerun.
