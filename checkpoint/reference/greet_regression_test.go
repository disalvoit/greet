// Reference receipt: red on main, green on the fix.
package greet

import "testing"

func TestHelloEmptyNameIsWorld(t *testing.T) {
	if got, want := Hello(""), "Hello, world!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", "", got, want)
	}
}

func TestHelloTrimsSurroundingSpace(t *testing.T) {
	if got, want := Hello("  Ada  "), "Hello, Ada!"; got != want {
		t.Errorf("Hello(%q) = %q, want %q", "  Ada  ", got, want)
	}
}
