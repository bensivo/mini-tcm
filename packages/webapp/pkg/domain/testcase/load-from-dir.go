package testcase

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func LoadFromDir(dirpath string) ([]TestCase, error) {
	tcs := []TestCase{}

	filepath.Abs(dirpath)

	// Get filepaths of all .tcm.md files in the directory (including subdirectories)
	filepaths, err := findFilesEndingWith(dirpath, `.tcm.md$`)
	if err != nil {
		return nil, err
	}

	// Read and parse the contents of each file
	for _, filepath := range filepaths {
		fmt.Printf("Parsing .tcm.md file: %s\n", filepath)

		bytes, err := os.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		str := string(bytes)

		tc, err := FromMarkdown(str)
		if err != nil {
			return nil, err
		}

		tcs = append(tcs, tc)
	}

	return tcs, nil
}

// findFilesEndingWith finds all files in a directory that end with a given suffix.
func findFilesEndingWith(dirpath string, suffixRegExp string) ([]string, error) {
	filepaths := []string{}
	err := filepath.WalkDir(dirpath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		matched, err := regexp.MatchString(suffixRegExp, path)
		if err != nil {
			return nil
		}
		if matched {
			filepaths = append(filepaths, path)
		}
		return nil
	})
	return filepaths, err
}
