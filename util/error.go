package util

import "fmt"

func NewError(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}

func NewErrorf(message string, args ...any) map[string]any {
	return map[string]any{
		"message": fmt.Sprintf(message, args...),
	}
}
