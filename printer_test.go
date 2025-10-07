package main

import (
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple message",
			input:    "Hello World",
			expected: "Message: Hello World",
		},
		{
			name:     "Empty message",
			input:    "",
			expected: "Message: ",
		},
		{
			name:     "Special characters",
			input:    "Hello @#$%",
			expected: "Message: Hello @#$%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Print(tt.input)
			if result != tt.expected {
				t.Errorf("Print(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPrintWithPrefix(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		msg      string
		expected string
	}{
		{
			name:     "With prefix",
			prefix:   "INFO",
			msg:      "System started",
			expected: "[INFO] System started",
		},
		{
			name:     "Empty prefix",
			prefix:   "",
			msg:      "test",
			expected: "[] test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrintWithPrefix(tt.prefix, tt.msg)
			if result != tt.expected {
				t.Errorf("PrintWithPrefix(%q, %q) = %q, want %q",
					tt.prefix, tt.msg, result, tt.expected)
			}
		})
	}
}

func TestPrintNumbers(t *testing.T) {
	result := PrintNumbers(1, 2, 3, 4, 5)
	expected := "Numbers: [1 2 3 4 5]"

	if result != expected {
		t.Errorf("PrintNumbers(1,2,3,4,5) = %q, want %q", result, expected)
	}
}

// Пример benchmark теста
func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Print("benchmark test")
	}
}
