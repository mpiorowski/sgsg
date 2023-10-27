package utils

import (
	"os"
	"testing"
)

func TestMustHaveEnv(t *testing.T) {
	// Test case 1: Environment variable exists
	os.Setenv("GOPATH", "/home/user/go")
	key := "GOPATH"
	value := mustHaveEnv(key)
	if value != "/home/user/go" {
		t.Errorf("Expected non-empty value for environment variable %v", key)
	}

	// Test case 2: Environment variable does not exist
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for non-existent environment variable")
		}
	}()
	os.Setenv("GOPATH", "")
	key = "GOPATH"
	_ = mustHaveEnv(key)
}
