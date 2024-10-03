package persianPhoneValidator

import (
	"testing"
)

func TestCleanPhoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0912-123-4567", "09121234567"},
		{"+98 912 123 4567", "989121234567"},
		{"(0912) 123 4567", "09121234567"},
		{"0098-912-123-4567", "00989121234567"},
		{"HelloWorld09121234567", "09121234567"},
		{"", ""}, // Test empty input
	}

	for _, test := range tests {
		result := CleanPhoneNumber(test.input)
		if result != test.expected {
			t.Errorf("CleanPhoneNumber(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"+989121234567", "09121234567"},
		{"00989121234567", "09121234567"},
		{"98121234567", "0121234567"},
		{"09121234567", "09121234567"},
		{"9121234567", "09121234567"}, // Missing leading zero
		{"12345", "012345"},           // Short input case
		{"", ""},                      // Test empty input
	}

	for _, test := range tests {
		result := NormalizePhoneNumber(test.input)
		if result != test.expected {
			t.Errorf("NormalizePhoneNumber(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestIsPhoneNumberValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"09121234567", true},    // Valid local format
		{"+989121234567", true},  // Valid international format
		{"00989121234567", true}, // Valid international format (0098)
		{"0912123456", false},    // Too short
		{"091212345678", false},  // Too long
		{"98121234567", false},
		{"9121234567", true},
		{"", false},            // Empty string
		{"abcdefghijk", false}, // Non-numeric input
		{"12345", false},       // Too short input
	}

	for _, test := range tests {
		result := IsPhoneNumberValid(test.input)
		if result != test.expected {
			t.Errorf("IsPhoneNumberValid(%s) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		input              string
		expectedValid      bool
		expectedNormalized string
	}{
		{"+989121234567", true, "09121234567"},
		{"00989121234567", true, "09121234567"},
		{"09121234567", true, "09121234567"},
		{"9121234567", true, "09121234567"},
		{"12345", false, "012345"},          // Invalid and normalized
		{"abcdefghijk", false, ""},          // Non-numeric invalid
		{"", false, ""},                     // Empty input
		{"0912123456", false, "0912123456"}, // Too short, invalid
	}

	for _, test := range tests {
		valid, normalized := ValidatePhoneNumber(test.input)
		if valid != test.expectedValid || normalized != test.expectedNormalized {
			t.Errorf("ValidatePhoneNumber(%s) = (%v, %s); expected (%v, %s)", test.input, valid, normalized, test.expectedValid, test.expectedNormalized)
		}
	}
}
