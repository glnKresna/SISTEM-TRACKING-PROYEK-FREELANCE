package utils

import "fmt"

// GenerateID generates a new ID for a project
func GenerateID(count int) string {
	return fmt.Sprintf("%03d", count+1)
}

// TruncateStr truncates a string to the specified length
func TruncateStr(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}
