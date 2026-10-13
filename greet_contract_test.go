package greet

import (
	"strings"
	"testing"
)

func TestHelloTrimmed(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"  Ada  ", "Hello, Ada!"},
		{"\tGrace\n", "Hello, Grace!"},
		{" \u00A0 Linus \u00A0 ", "Hello, Linus!"},
	}
	for _, tc := range cases {
		if got := Hello(tc.name); got != tc.want {
			t.Errorf("Hello(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestHelloEmptyAfterTrim(t *testing.T) {
	cases := []string{"", "   ", "\t\n", "\u00A0\u00A0"}
	for _, name := range cases {
		if got := Hello(name); got != "Hello, world!" {
			t.Errorf("Hello(%q) = %q, want %q", name, got, "Hello, world!")
		}
	}
}

func TestHelloInternalWhitespacePreserved(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"Ada  Lovelace", "Hello, Ada  Lovelace!"},
		{"Ada\tLovelace", "Hello, Ada\tLovelace!"},
		{"Ada\nLovelace", "Hello, Ada\nLovelace!"},
	}
	for _, tc := range cases {
		if got := Hello(tc.name); got != tc.want {
			t.Errorf("Hello(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestHelloCasePreserved(t *testing.T) {
	if got := Hello("aDa"); got != "Hello, aDa!" {
		t.Errorf("Hello(%q) = %q, want %q", "aDa", got, "Hello, aDa!")
	}
}

func TestHelloUnicodePreserved(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"👋 Ada", "Hello, 👋 Ada!"},
		{"José", "Hello, José!"},
		{"山田", "Hello, 山田!"},
	}
	for _, tc := range cases {
		if got := Hello(tc.name); got != tc.want {
			t.Errorf("Hello(%q) = %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestHelloNoPanic(t *testing.T) {
	// Empty
	Hello("")
	// Huge (10MB)
	Hello(strings.Repeat("x", 10_000_000))
	// NUL bytes
	Hello("Ada\000Lovelace")
	// Invalid UTF-8
	Hello("\xff\xfe")
}

func TestHelloExactFormat(t *testing.T) {
	// No extra spaces, exact format
	if got := Hello("Ada"); got != "Hello, Ada!" {
		t.Errorf("exact format: got %q", got)
	}
	// Empty after trim
	if got := Hello(" "); got != "Hello, world!" {
		t.Errorf("empty after trim: got %q", got)
	}
}