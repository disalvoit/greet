// Held-out core suite for the greet checkpoint.
//
// This file is NOT part of the repository the agent works in: the grader copies
// it into the module root at grade time, runs `go test -v -run TestCore .`, and
// then removes it. It is an external test package, so it cannot collide with the
// package's own tests.
package greet_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/disalvoit/greet"
)

// heldOutCLI runs `go run ./cmd/greet <args...>` from the module root and
// returns stdout, stderr and the exit code. Defined once here and shared with
// the edge suite, which is always copied in alongside this file.
func heldOutCLI(t *testing.T, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command("go", append([]string{"run", "./cmd/greet"}, args...)...)
	var out, errb strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &errb
	err := cmd.Run()
	code := 0
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("go run ./cmd/greet %v failed to start: %v", args, err)
		}
		code = ee.ExitCode()
	}
	return out.String(), errb.String(), code
}

func eq(t *testing.T, got, want, label string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %q, want %q", label, got, want)
	}
}

func TestCoreSimple(t *testing.T) {
	eq(t, greet.Hello("Ada"), "Hello, Ada!", `Hello("Ada")`)
}

func TestCoreEmpty(t *testing.T) {
	eq(t, greet.Hello(""), "Hello, world!", `Hello("")`)
}

func TestCoreSpacesOnly(t *testing.T) {
	eq(t, greet.Hello("   "), "Hello, world!", `Hello("   ")`)
}

func TestCoreTrimLeading(t *testing.T) {
	eq(t, greet.Hello("  Ada"), "Hello, Ada!", `Hello("  Ada")`)
}

func TestCoreTrimTrailing(t *testing.T) {
	eq(t, greet.Hello("Ada  "), "Hello, Ada!", `Hello("Ada  ")`)
}

func TestCoreTrimTabsNewlines(t *testing.T) {
	eq(t, greet.Hello("\t Ada \n"), "Hello, Ada!", `Hello("\t Ada \n")`)
}

func TestCoreInternalSpaceKept(t *testing.T) {
	eq(t, greet.Hello("Ada  Lovelace"), "Hello, Ada  Lovelace!", `Hello("Ada  Lovelace")`)
}

func TestCoreCasePreserved(t *testing.T) {
	eq(t, greet.Hello("aDA lovelace"), "Hello, aDA lovelace!", `Hello("aDA lovelace")`)
}

func TestCoreExactFormat(t *testing.T) {
	eq(t, greet.Hello("Ada!"), "Hello, Ada!!", `Hello("Ada!")`)
	eq(t, greet.Hello("Hello, Ada!"), "Hello, Hello, Ada!!", `Hello("Hello, Ada!")`)
}

func TestCoreUnicodeName(t *testing.T) {
	eq(t, greet.Hello("Ágata 👋"), "Hello, Ágata 👋!", `Hello("Ágata 👋")`)
}

func TestCoreNeverPanics(t *testing.T) {
	for _, name := range []string{"", "\x00", "\xff\xfe", "\u00a0", strings.Repeat("a", 1<<20)} {
		got := greet.Hello(name)
		if !strings.HasPrefix(got, "Hello, ") || !strings.HasSuffix(got, "!") {
			t.Errorf("Hello(%.20q…) = %.40q; want the shape `Hello, …!`", name, got)
		}
	}
}

func TestCoreWorldIsNotSpecial(t *testing.T) {
	eq(t, greet.Hello("world"), "Hello, world!", `Hello("world")`)
	eq(t, greet.Hello(" world "), "Hello, world!", `Hello(" world ")`)
}
