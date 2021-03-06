package urlvalidator

import "net/url"

// IsValidURL checks if a given URL is valid
func IsValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Host == "" {
		return false
	}

	if u.Scheme != "" && u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}
