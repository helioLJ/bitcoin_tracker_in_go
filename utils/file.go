package utils

import (
	"os"
)

func AppendToFile(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return WrapError("failed to open file", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return WrapError("failed to write to file", err)
	}
	return nil
}
