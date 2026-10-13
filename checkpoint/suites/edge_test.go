// Held-out edge suite for the greet checkpoint. Copied into the module root at
// grade time, run as `go test -v -run TestEdge .`, then removed. It shares the
// heldOutCLI/eq helpers with core_test.go, which is always copied in with it.
package greet_test

import (
	"strings"
	"testing"

	"github.com/disalvoit/greet"
)

func TestEdgeLongName(t *testing.T) {
	name := strings.Repeat("Ada", 4000) // 12,000 characters
	eq(t, greet.Hello(name), "Hello, "+name+"!", "Hello(12k chars)")
}

func TestEdgeInvalidUTF8(t *testing.T) {
	name := "\xff\xfe"
	eq(t, greet.Hello(name), "Hello, \xff\xfe!", `Hello("\xff\xfe")`)
}

func TestEdgeNulByte(t *testing.T) {
	name := "a\x00b"
	eq(t, greet.Hello(name), "Hello, a\x00b!", `Hello("a\x00b")`)
}

func TestEdgeNewlineInside(t *testing.T) {
	eq(t, greet.Hello("Ada\nLovelace"), "Hello, Ada\nLovelace!", `Hello("Ada\nLovelace")`)
}

func TestEdgeUnicodeSpace(t *testing.T) {
	// U+00A0 is whitespace to strings.TrimSpace, so it must not survive.
	eq(t, greet.Hello("\u00a0Ada\u00a0"), "Hello, Ada!", `Hello("\u00a0Ada\u00a0")`)
}

func TestEdgeEmoji(t *testing.T) {
	eq(t, greet.Hello(" 👋 Ada 👋 "), "Hello, 👋 Ada 👋!", `Hello(" 👋 Ada 👋 ")`)
}

func TestEdgeCRLF(t *testing.T) {
	eq(t, greet.Hello("Ada\r\n"), "Hello, Ada!", `Hello("Ada\r\n")`)
}

func TestEdgeTabOnly(t *testing.T) {
	eq(t, greet.Hello("\t\t"), "Hello, world!", `Hello("\t\t")`)
}

func TestEdgeCLIDefault(t *testing.T) {
	out, errOut, code := heldOutCLI(t)
	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if out != "Hello, world!\n" {
		t.Errorf("stdout = %q, want %q", out, "Hello, world!\n")
	}
	if errOut != "" {
		t.Errorf("stderr = %q, want empty", errOut)
	}
}

func TestEdgeCLIFirstArgOnly(t *testing.T) {
	out, errOut, code := heldOutCLI(t, "Ada", "Lovelace")
	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if out != "Hello, Ada!\n" {
		t.Errorf("stdout = %q, want %q", out, "Hello, Ada!\n")
	}
	if errOut != "" {
		t.Errorf("stderr = %q, want empty", errOut)
	}
}
