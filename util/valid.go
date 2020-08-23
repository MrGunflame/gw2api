package util

import "regexp"

// ValidAPIKey checks whether the entered apikey has a valid format
func ValidAPIKey(apikey string) bool {
	return regexp.MustCompile(`([A-Z0-9]{8}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{20}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{12})`).MatchString(apikey)
}
