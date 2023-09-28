package test

import (
	"go/token"
	"testing"
)

func TestIsKeyword(t *testing.T) {
	testCases := []struct {
		tokenType string
		expected  bool
	}{
		{"for", true},      // Expected: true, as FOR is a keyword
		{"ident", false},   // Expected: false, as IDENT is not a keyword
		{"break", true},    // Expected: true, as BREAK is a keyword
		{"string", false},  // Expected: false, as STRING is not a keyword
		{"import", true},   // Expected: true, as IMPORT is a keyword
		{"add", false},     // Expected: false, as ADD is not a keyword
		{"illegal", false}, // Expected: false, as ILLEGAL is not a keyword
	}

	for _, tc := range testCases {
		actual := token.IsKeyword(tc.tokenType)
		if actual != tc.expected {
			t.Errorf("Unexpected result for token type %v: got %v, want %v", tc.tokenType, actual, tc.expected)
		}
	}
}
