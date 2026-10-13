package greet

import (
	"strings"
	"testing"
)

func TestHelloUsesWorldForEmptyName(t *testing.T) {
	for _, n := range []string{"", "   ", "\t\n", "\u00a0"} {
		if got, want := Hello(n), "Hello, world!"; got != want {
			t.Errorf("Hello(%q) = %q, want %q", n, got, want)
		}
	}
}

func TestHelloTrimsSurroundingWhitespace(t *testing.T) {
	cases := map[string]string{
		"  Ada  ":      "Hello, Ada!",
		"\tAda\n":      "Hello, Ada!",
		"\u00a0Grace ": "Hello, Grace!",
	}
	for in, want := range cases {
		if got := Hello(in); got != want {
			t.Errorf("Hello(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestHelloPreservesInternalWhitespaceAndCase(t *testing.T) {
	const in = "  Ada  LOVELACE  "
	if got, want := Hello(in), "Hello, Ada  LOVELACE!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", in, got, want)
	}
}

func TestHelloDoesNotPanicOnArbitraryInput(t *testing.T) {
	inputs := []string{"", strings.Repeat("x", 1<<20), "a\x00b", string([]byte{0xff, 0xfe})}
	for _, in := range inputs {
		_ = Hello(in)
	}
}
