package validator

import (
	nu "net/url"
	"regexp"
	"strings"
)

var validDomains = [...]string{
	"com",
	"net",
	"org",
	"io",
	"vercel",
	"app",
}

func ValidateURL(url string) bool {
	hasSpacesRex := regexp.MustCompile(`\s+`)

	// Basic check for spaces
	if hasSpacesRex.MatchString(url) {
		return false
	}

	// Parse the URL and check the scheme and domain
	parsed, err := nu.Parse(url)
	if err != nil {
		return false
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	host := parsed.Host
	// Remove port if present
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return false
	}
	domain := parts[len(parts)-1]

	for _, valid := range validDomains {
		if domain == valid {
			return true
		}
	}
	return false
}