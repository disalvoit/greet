# greet — the contract

A tiny, dependency-free Go module. This file is the contract: the held-out suite
grades `greet` against exactly what is written here, and nothing else.

## Library

Module path: `github.com/disalvoit/greet`. Package `greet` lives at the module root.

    func Hello(name string) string

Rules, in order:

1. Leading and trailing whitespace is trimmed from `name`, with Go's
   `strings.TrimSpace` semantics (spaces, tabs, newlines, carriage returns and
   other Unicode space characters such as U+00A0).
2. If the trimmed name is empty — `""`, `"   "`, `"\t\n"` — the greeting is
   `Hello, world!` exactly. The word `world` is not special in any other way:
   `Hello("world")` is `Hello, world!` only because `world` is the name.
3. Otherwise the greeting is `Hello, ` + trimmed name + `!`. The trimmed name is
   used **verbatim**: no case changes, no escaping, no ASCII filtering, no
   collapsing of internal whitespace. `Hello("Ada Lovelace")` is
   `Hello, Ada Lovelace!` and `Hello("👋 Ada")` is `Hello, 👋 Ada!`.
4. `Hello` never panics, for any string: empty, megabytes long, containing NUL
   bytes or invalid UTF-8.

## Command line

    go run ./cmd/greet              -> Hello, world!   (exit 0)
    go run ./cmd/greet Ada          -> Hello, Ada!     (exit 0)
    go run ./cmd/greet "  Ada  "    -> Hello, Ada!     (exit 0)
    go run ./cmd/greet Ada Lovelace -> Hello, Ada!     (exit 0, first argument only)

The greeting goes to stdout followed by a single newline; nothing is written to
stderr.

## Module hygiene

* **Standard library only.** `go.mod` declares no `require` dependencies, and the
  module builds and tests with `GOPROXY=off` (no network, no module cache).
* `go vet ./...` is clean.
* `go test ./...` **discovers and runs** tests for the `greet` package. If that
  package reports `[no test files]`, the module has no test coverage where it
  matters — that is a failure. So is a test that passes without asserting
  anything: an empty table, a loop over zero cases, or a test that only calls
  `t.Skip`. (Packages that legitimately hold no tests, such as `cmd/greet`, are
  not counted against you.)
