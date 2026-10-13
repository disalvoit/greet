// Reference fix for the greet checkpoint (the answer, not the task).
package greet

import "strings"

// Hello returns a greeting for name, trimming surrounding whitespace and
// falling back to "world" when nothing is left.
func Hello(name string) string {
	if n := strings.TrimSpace(name); n != "" {
		return "Hello, " + n + "!"
	}
	return "Hello, world!"
}
