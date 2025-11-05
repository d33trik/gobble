package main_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binName = "gobble_test"

func TestMain(m *testing.M) {
	cmd := exec.Command("go", "build", "-o", binName, ".")

	errBuf := &bytes.Buffer{}
	cmd.Stderr = errBuf

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to build binary:", err, errBuf.String())
		os.Exit(1)
	}

	result := m.Run()

	os.Remove(binName)
	os.Exit(result)
}

func TestStdin(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("failed to get working directory:", err)
	}

	path := filepath.Join(dir, binName)

	cmd := exec.Command(path)

	output := &bytes.Buffer{}

	cmd.Stdin = strings.NewReader("one two three four five six\n")
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	got := output.String()
	want := " 1 6 28\n"

	if got != want {
		t.Logf("got: %v, want: %v", got, want)
		t.Fail()
	}
}
