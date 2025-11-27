package xvalidator

import "strings"

// trimStructName removes the struct name prefix if it exists
func trimStructName(field string) string {
	// Remove the struct name prefix if it exists
	if idx := strings.Index(field, "."); idx != -1 {
		return field[idx+1:]
	}
	return field
}
