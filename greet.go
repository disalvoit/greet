// Package greet turns a name into a greeting.
package greet

import "strings"

// Hello returns a greeting for name.
func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, world!"
	}
	return "Hello, " + name + "!"
}
