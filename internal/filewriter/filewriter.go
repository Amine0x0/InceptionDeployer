package filewriter

import (
	"os"
	"path/filepath"
)

func Write(path string, content string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), mode)
}
