#!/bin/sh
# Episode 01 checkpoint — 22 held-out tests against YOUR greet module.
#
#   ./run-checkpoint.sh              # grade the module in this directory
#   ./run-checkpoint.sh ~/src/greet  # grade a module elsewhere
#
# Your working tree is never modified: the module is copied to a temp dir first.
set -eu

HERE=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
MODULE=${1:-$PWD}
IMAGE=${CHECKPOINT_IMAGE:-golang:1.24-alpine}

WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT

# 1. copy the module you are grading (no .git, no node_modules, no checkpoint dir)
tar -cf - -C "$MODULE" --exclude=.git --exclude=checkpoint . 2>/dev/null | tar -xf - -C "$WORK"

# 2. drop the held-out suites in next to the package under test
cp "$HERE/suites/core_test.go" "$HERE/suites/edge_test.go" "$WORK/"

# 3. run them with the pinned toolchain, with the network switched off
docker run --rm \
  -e GOPROXY=off -e GOTOOLCHAIN=local -e GOFLAGS=-mod=mod \
  -v "$WORK":/w -w /w "$IMAGE" sh -lc '
    export PATH=/usr/local/go/bin:$PATH
    go vet ./... && echo "vet: clean" || echo "vet: FAILED"
    go test -v ./... > /tmp/out.txt 2>&1 || true
    core_pass=$(grep -c "^--- PASS: TestCore" /tmp/out.txt || true)
    core_fail=$(grep -c "^--- FAIL: TestCore" /tmp/out.txt || true)
    edge_pass=$(grep -c "^--- PASS: TestEdge" /tmp/out.txt || true)
    edge_fail=$(grep -c "^--- FAIL: TestEdge" /tmp/out.txt || true)
    total=$((core_pass + core_fail + edge_pass + edge_fail))
    echo
    echo "Episode 01 checkpoint - held-out suite"
    echo "  core  $core_pass / $((core_pass + core_fail))"
    echo "  edge  $edge_pass / $((edge_pass + edge_fail))"
    echo "  total $((core_pass + edge_pass)) / $total passed"
    if [ "$core_fail" -gt 0 ] || [ "$edge_fail" -gt 0 ]; then
      echo
      echo "failing:"
      grep "^--- FAIL" /tmp/out.txt | sed "s/--- FAIL: /  /" || true
      exit 1
    fi
    echo
    echo "all held-out checks pass."
  '
