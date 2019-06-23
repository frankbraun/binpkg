package pkg

import (
	"fmt"
	"os"
)

// dirExists checks if filename exists and is a directory
func dirExists(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("'%s' is not a directory", filename)
	}
	return nil
}
