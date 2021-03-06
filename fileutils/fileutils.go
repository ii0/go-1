package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CopyDirectory copies a directory from src to dst recursively.
func CopyDirectory(src, dst string) error {
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcFileInfo.Mode())
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if !entry.IsDir() {
			cont, err := ioutil.ReadFile(srcPath)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dstPath, cont, entry.Mode())
			if err != nil {
				return err
			}
		} else {
			err = CopyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
