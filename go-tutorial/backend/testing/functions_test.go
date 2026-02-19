package main

import "testing"

// Testing generic functions

// Simple Unit Test
func TestAddNumbers(t *testing.T) {
	result := addNumbers(10, 35)
	expected := 45

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// Table-Driven test with multiple cases
func TestValidateUsername(t *testing.T) {

    // 1. Define the "Table" (a slice of structs)
    tests := []struct {
        name     string // Description of the case
        input    string // The value to test
        expected bool   // What we want to happen
    }{
        {"valid username", "gopher123", true},
        {"too short", "go", false},
        {"too long", "thisusernameiswaytoolongforoursystem", false},
        {"empty string", "", false},
    }

    // 2. Iterate over the table
    for _, tt := range tests {
        // 3. Use t.Run to create a "Subtest"
        t.Run(tt.name, func(t *testing.T) {
            got := ValidateUsername(tt.input)
            if got != tt.expected {
                t.Errorf("ValidateUsername(%q) = %v; want %v", tt.input, got, tt.expected)
            }
        })
    }
}

/*
Why use t.Run?

By using t.Run, each case becomes a subtest. This allows you to:

    Identify failures easily: If "too short" fails, the terminal shows TestValidateUsername/too_short failed, while the others passed.

    Run one specific case: go test -v -run=TestValidateUsername/special_characters.
*/