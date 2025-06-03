package utils

import "fmt"

// Generate ID untuk setiap data proyek terbaru
func GenerateID(count int) string {
	return fmt.Sprintf("%03d", count+1)
}

// Truncate (potong) string ke suatu panjang tertentu
func TruncateStr(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}
