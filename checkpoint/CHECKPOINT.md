# Episode 01 — the checkpoint

Twenty-two held-out tests for the `greet` module, graded against the contract in the
repository's `README.md` and against nothing else. They are not in the module's own
test suite: a checkpoint you can read before you write the code would not be a
checkpoint.

## Run it

    git clone https://github.com/disalvoit/greet          # the starter
    git clone -b checkpoint https://github.com/disalvoit/greet checkpoint-suite
    ./checkpoint-suite/run-checkpoint.sh ./greet

Needs `docker` and nothing else: the tests run inside the pinned toolchain from the
episode, `golang:1.24-alpine`, with `GOPROXY=off` — no network, no module cache, no
dependencies. Your working tree is copied to a temporary directory first, so the run
never modifies your code.

Output on the starter, measured:

    core  6 / 12
    edge  5 / 10
    total 11 / 22 passed

`reference/greet.go` is the fix from the episode. Do not read it before you have
written a test that fails without it — that is the habit the episode is about.

## What it grades

Twelve core tests (`suites/core_test.go`): trimming, the empty-name fallback,
case and internal spacing preserved, `world` not being special.

Ten edge tests (`suites/edge_test.go`): Unicode spaces, emoji, CRLF, tabs, very long
names, and the command line itself.

Both files are `package greet` and are copied next to the package under test at run
time. They import nothing outside the standard library.
