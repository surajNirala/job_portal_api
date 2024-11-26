package utils

import (
	"fmt"
	"os"
)

func DeleteFileIfExists(path string) error {
	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// File does not exist, return without error
		return nil
	}

	// Try to delete the file
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}
	return nil
}
