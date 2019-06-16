package util

import "strings"

// SanitizeURL removes unnecssary data from the CustomURLs
func SanitizeURL(str string) string {
	str = strings.ReplaceAll(str, " ", "_")
	return str
}
