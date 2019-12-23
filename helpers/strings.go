package helpers

import "strings"

// Convert string to lower case and trim all whitespaces
func ToLowerAndTrim(value string) string {
    return strings.TrimSpace(strings.ToLower(value))
}

// Convert string to lower case and replace given values
func ToLowerAndReplace(value string, replace string, with string) string {
    return ToLowerAndTrim(strings.ReplaceAll(value, replace, with))
}