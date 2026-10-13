package greet

import (
	"strings"
	"testing"
)

func TestHelloTrimsSurroundingWhitespace(t *testing.T) {
	cases := []struct{ name, in, want string }{
		{"spaces", "  Ada  ", "Hello, Ada!"},
		{"tabs", "\tAda\t", "Hello, Ada!"},
		{"newlines", "\nAda\n", "Hello, Ada!"},
		{"carriage returns", "\rAda\r", "Hello, Ada!"},
		{"unicode nbsp", "\u00a0Ada\u00a0", "Hello, Ada!"},
		{"mixed", " \t\r\n Ada \t\r\n ", "Hello, Ada!"},
	}
	for _, c := range cases {
		if got := Hello(c.in); got != c.want {
			t.Errorf("Hello(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHelloEmptyNameGreetsWorld(t *testing.T) {
	for _, in := range []string{"", "   ", "\t\n", " \u00a0 "} {
		if got, want := Hello(in), "Hello, world!"; got != want {
			t.Errorf("Hello(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestHelloWorldIsJustAName(t *testing.T) {
	if got, want := Hello("world"), "Hello, world!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", "world", got, want)
	}
	if got, want := Hello("World"), "Hello, World!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", "World", got, want)
	}
	if got, want := Hello("  world  "), "Hello, world!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", "  world  ", got, want)
	}
}

func TestHelloPreservesNameVerbatim(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Ada Lovelace", "Hello, Ada Lovelace!"},
		{"  Ada Lovelace  ", "Hello, Ada Lovelace!"},
		{"Ada   Lovelace", "Hello, Ada   Lovelace!"},
		{"a\tb\nc", "Hello, a\tb\nc!"},
		{"gRaCe hOpPeR", "Hello, gRaCe hOpPeR!"},
		{"  spaced  inner  ", "Hello, spaced  inner!"},
	}
	for _, c := range cases {
		if got := Hello(c.in); got != c.want {
			t.Errorf("Hello(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHelloUnicodeNamesPreserved(t *testing.T) {
	cases := []struct{ in, want string }{
		{"👋 Ada", "Hello, 👋 Ada!"},
		{"  👋 Ada  ", "Hello, 👋 Ada!"},
		{"Ada Lovelace 🇬🇧", "Hello, Ada Lovelace 🇬🇧!"},
		{"Ωμέγα", "Hello, Ωμέγα!"},
		{"日本語", "Hello, 日本語!"},
	}
	for _, c := range cases {
		if got := Hello(c.in); got != c.want {
			t.Errorf("Hello(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHelloNeverPanics(t *testing.T) {
	inputs := []string{
		"",
		" ",
		strings.Repeat("x", 5<<20),
		strings.Repeat(" ", 1<<20),
		"a\x00b",
		"\x00",
		"nul\x00 \x00 ",
		"\xff",
		"\xed\xa0\x80",
		"  \xff\xfe  ",
	}
	for _, in := range inputs {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Hello(%q...) panicked: %v", truncate(in), r)
				}
			}()
			got := Hello(in)
			if !strings.HasPrefix(got, "Hello, ") || !strings.HasSuffix(got, "!") {
				t.Errorf("Hello(%q...) = %q, want form \"Hello, <name>!\"", truncate(in), got)
			}
		}()
	}
}

func truncate(s string) string {
	if len(s) > 32 {
		return s[:32] + "..."
	}
	return s
}
