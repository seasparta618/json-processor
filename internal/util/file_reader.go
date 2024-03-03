package util

import (
	"fmt"
	apperrors "json-processor/internal/error"
	"os"
)

// ReadFile reads the contents of the specified file and returns them as a string.
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", apperrors.ErrorFileNotFound
		}
		return "", fmt.Errorf("unexpected error reading file: %w", err)
	}
	if len(data) == 0 {
		return "", apperrors.ErrorEmptyContent
	}
	return string(data), nil
}
