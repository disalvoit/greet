# Measured on main

Everything below is output from this repository at commit `b45d1b1`
("greet: hello + one green test"), run inside the pinned toolchain container
`golang:1.24-alpine` (`go version go1.24.13 linux/arm64`) with `GOPROXY=off`
and `GOTOOLCHAIN=local`, so nothing here can reach the network:

    docker run --rm -e GOPROXY=off -e GOFLAGS=-mod=mod \
      -v gobuild:/root/.cache -v "$PWD":/w -w /w golang:1.24-alpine \
      sh -c 'export PATH=/usr/local/go/bin:$PATH; go vet ./... && go test ./...'

## The public signal is green

    $ go vet ./...
    (clean)

    $ go test -v ./...
    === RUN   TestHelloKnownNames
    --- PASS: TestHelloKnownNames (0.00s)
    PASS
    ok      github.com/disalvo/greet        0.005s
    ?       github.com/disalvo/greet/cmd/greet      [no test files]

One test. It walks three names — `Ada`, `Grace`, `Linus` — and all three are
non-empty and whitespace-free.

## The program is wrong

    $ go run ./cmd/greet
    Hello, !

    $ go run ./cmd/greet "  Ada  "
    Hello,   Ada  !

## The held-out suite, against main

| suite | pass | fail |
| --- | --- | --- |
| core (12) | 6 | 6 |
| edge (10) | 5 | 5 |

Core failures: `TestCoreEmpty`, `TestCoreSpacesOnly`, `TestCoreTrimLeading`,
`TestCoreTrimTrailing`, `TestCoreTrimTabsNewlines`, `TestCoreWorldIsNotSpecial`.

Edge failures: `TestEdgeUnicodeSpace`, `TestEdgeEmoji`, `TestEdgeCRLF`,
`TestEdgeTabOnly`, `TestEdgeCLIDefault`.

**11 of the 22 held-out checks fail on a branch where `go test ./...` is green
and CI is passing.** The suite is not lying about what it ran; it is lying about
what it covered.
