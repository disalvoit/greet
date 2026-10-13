package greet

import "testing"

func TestHelloKnownNames(t *testing.T) {
	for _, n := range []string{"Ada", "Grace", "Linus"} {
		if got, want := Hello(n), "Hello, "+n+"!"; got != want {
			t.Errorf("Hello(%q) = %q, want %q", n, got, want)
		}
	}
}
