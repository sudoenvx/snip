package validator

import "testing"

type TestUrl struct {
	url      string
	isValid bool
}

func TestValidateUrl(t *testing.T) {
	testUrls := []TestUrl{
		// Valid URLs
		{"https://www.google.com", true},
		{"http://example.net", true},
		{"https://sub.domain.org/path", true},
		{"http://localhost.io", true},
		{"https://example.com/path/to/resource?query=1", true},
		// Invalid URLs
		{"", false},
		{"ftp://example.com", false},
		{"http:/example.com", false},
		{"https://", false},
		{"https://example", false},
		{"https://example.c", false},
		{"https://example.123", false},
		{"https://example.com some text", false},
		{"https://example.com\t", false},
		{"https:// example.com", false},
		{"www.example.com", false},
		{"example.com", false},
	}

	for _, test := range testUrls {
		result := ValidateURL(test.url)
		if result != test.isValid {
			t.Errorf("ValidateURL(%q) = %v; want %v", test.url, result, test.isValid)
		}
	}
}
