package utils

import (
	"strings"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type TestStruct struct {
		Name        string `validate:"required"`
		Description string `validate:"required,min=10,max=100"`
	}

	// Test case 1: Struct is valid
	testStruct := TestStruct{Name: "test", Description: "test description"}
	rules := map[string]string{"TestStruct": "required"}
	err := ValidateStruct(rules, testStruct, &testStruct)
	if err != nil {
		t.Errorf("Error validating struct: %v", err)
	}

	// Test case 2: Struct is invalid
	testStruct = TestStruct{Name: ""}
	err = ValidateStruct(rules, testStruct, &testStruct)
	containsName := strings.Contains(err.Error(), "Name") && strings.Contains(err.Error(), "required") && strings.Contains(err.Error(), "InvalidArgument")
	containsDescription := strings.Contains(err.Error(), "Description") && strings.Contains(err.Error(), "required") && strings.Contains(err.Error(), "InvalidArgument")
	if !containsName || !containsDescription {
		t.Errorf("Error validating struct: %v", err)
	}

    // Test case 3: Struct is invalid
    testStruct = TestStruct{Name: "test", Description: "test"}
    err = ValidateStruct(rules, testStruct, &testStruct)
    containsDescription = strings.Contains(err.Error(), "Description") && strings.Contains(err.Error(), "min") && strings.Contains(err.Error(), "InvalidArgument")
    if !containsDescription {
        t.Errorf("Error validating struct: %v", err)
    }
}
