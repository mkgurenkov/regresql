package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

// Check that given path string is a directory that exists
func checkDirectory(cwd string) error {
	stat, err := os.Stat(cwd)
	if err != nil {
		return fmt.Errorf("No directory found at '%s'\n", cwd)
	}
	if !stat.IsDir() {
		return fmt.Errorf("Not a directory: '%s'\n", cwd)
	}
	return nil
}

func checkFiles(root string, files []string) error {
	if files == nil {
		return nil
	}

	for _, file := range files {
		stat, _ := os.Stat(file)

		if stat.IsDir() {
			return fmt.Errorf("Is a directory: '%s'\n", file)
		}

		if root != filepath.Dir(file) {
			return fmt.Errorf("'%s' parent directory does not match the root directory: '%s'\n", file, cwd)
		}
	}
	return nil
}

func expandPattern(pattern string) ([]string, error) {
	if pattern == "no" {
		return nil, nil
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("Bad patten %q: %v\n", pattern, err)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("No files found\n")
	}

	return matches, nil
}
