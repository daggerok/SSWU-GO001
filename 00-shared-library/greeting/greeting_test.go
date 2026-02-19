package greeting_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/00-shared-library/greeting"
)

func TestGreeting(t *testing.T) {
	// 1. Setup t cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Nil pointer",
			input:    "",
			expected: "Hello, Anonymous!",
		},
		{
			name:     "Empty string pointer",
			input:    " \n \t",
			expected: "Hello, Anonymous!",
		},
		{
			name:     "Valid name",
			input:    "Gopher",
			expected: "Hello, Gopher!",
		},
	}

	// 2. Run t cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := greeting.Greeting(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}
