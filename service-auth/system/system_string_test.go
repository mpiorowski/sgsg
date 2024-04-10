package system

import (
	"encoding/base64"
	"testing"
)

func TestContainsString(t *testing.T) {
	// Test case 1: String exists in the slice
	slice1 := []string{"apple", "banana", "cherry", "date"}
	str1 := "banana"
	if !ContainsString(slice1, str1) {
		t.Errorf("Expected true, but got false for %v in %v", str1, slice1)
	}

	// Test case 2: String does not exist in the slice
	slice2 := []string{"one", "two", "three", "four"}
	str2 := "five"
	if ContainsString(slice2, str2) {
		t.Errorf("Expected false, but got true for %v in %v", str2, slice2)
	}
}

func TestGenerateRandomState(t *testing.T) {
	// Test case: Generate a random state with a specific length
	length := 32
	state, err := GenerateRandomState(length)
    if err != nil {
        t.Errorf("Error generating random state: %v", err)
    }
	expectedLength := base64.StdEncoding.EncodedLen(length)
	if len(state) != expectedLength {
		t.Errorf("Expected state length of %d, but got %d", expectedLength, len(state))
	}
}

func TestGenerateRandomString(t *testing.T) {
	// Test case: Generate a random string with a specific length
	length := 16
	randomString, err := GenerateRandomString(length)
	if err != nil {
		t.Errorf("Error generating random string: %v", err)
	}
	if len(randomString) != length*2 { // Hex encoding doubles the length
		t.Errorf("Expected string length of %d, but got %d", length*2, len(randomString))
	}
}
