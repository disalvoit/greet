package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func buildGreet(t *testing.T) string {
	t.Helper()
	name := "greet"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	bin := filepath.Join(t.TempDir(), name)
	out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput()
	if err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

func TestCLIHonorsGreetingContract(t *testing.T) {
	bin := buildGreet(t)
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"no argument greets world", nil, "Hello, world!"},
		{"single name", []string{"Ada"}, "Hello, Ada!"},
		{"surrounding whitespace trimmed", []string{"  Ada  "}, "Hello, Ada!"},
		{"only first argument used", []string{"Ada", "Lovelace"}, "Hello, Ada!"},
		{"empty after trim greets world", []string{"\t\n"}, "Hello, world!"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cmd := exec.Command(bin, c.args...)
			var stdout, stderr strings.Builder
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				t.Fatalf("exit status: %v, stderr=%q", err, stderr.String())
			}
			if got := stderr.String(); got != "" {
				t.Errorf("stderr = %q, want empty", got)
			}
			if got, want := stdout.String(), c.want+"\n"; got != want {
				t.Errorf("stdout = %q, want %q", got, want)
			}
		})
	}
}
