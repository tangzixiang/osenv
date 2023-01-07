package main_test

import (
	"os"
	"runtime"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Logf("os.Environ():")
	for _, env := range os.Environ() {
		t.Log(env)
	}

	t.Logf("runtime.GOOS: %v", runtime.GOOS)
}
