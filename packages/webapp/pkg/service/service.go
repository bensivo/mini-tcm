package service

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bensivo/mini-tcm/packages/webapp/pkg/model"
	"gopkg.in/yaml.v3"
)

// LoadTestCasesFromDir loads test cases from files in a directory.
//
// It searches for any files with a '.tcm.md' extension, and parses the contents
func LoadTestCasesFromDir(dirpath string) ([]model.TestCase, error) {
	tcs := []model.TestCase{}

	filepath.Abs(dirpath)

	// Get filepaths of all .tcm.md files in the directory (including subdirectories)
	filepaths, err := FindFilesEndingWith(dirpath, `.tcm.md$`)
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

		tc, err := ParseTestCaseMd(str)
		if err != nil {
			return nil, err
		}

		tc.Print()
		tcs = append(tcs, tc)
	}

	return tcs, nil
}

// FindFilesEndingWith finds all files in a directory that end with a given suffix.
func FindFilesEndingWith(dirpath string, suffixRegExp string) ([]string, error) {
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

// ParseTestCaseMd parses a test case from a markdown string.
//
// We assume the markdown containins yaml front-matter (with the id and name fields),
// followed by sections for the spec, steps, and expected result.
//
// Each section is identified by a markdown header (e.g. '## Spec'), then the following lines
// until the next header are considered part of that section.
func ParseTestCaseMd(md string) (model.TestCase, error) {

	tc := model.TestCase{}

	// Extract the yaml front matter
	r := regexp.MustCompile(`(?s)---\n(.*?)\n---\n`)
	match := r.FindStringSubmatch(md)
	if len(match) < 2 {
		fmt.Println("no yaml front matter found")
		return tc, fmt.Errorf("no yaml front matter found")
	}
	frontMatterStr := match[1]

	var frontMatter struct {
		Id   string `yaml:"id"`
		Name string `yaml:"name"`
	}
	err := yaml.Unmarshal([]byte(frontMatterStr), &frontMatter)
	if err != nil {
		fmt.Println(err)
		return tc, err
	}

	tc.Id = frontMatter.Id
	tc.Name = frontMatter.Name

	r = regexp.MustCompile(`^#+\s`) // Regex to check for md headers
	sections := map[string]string{}
	var header string
	var content string
	lines := strings.Split(md, "\n")

	for _, line := range lines {
		match := r.MatchString(line) // See if we hit a new heading
		if err != nil {
			return tc, err
		}
		if match {
			if header != "" {
				sections[header] = content // Push the previous header + content into the map
			}

			header = strings.Trim(line, "# ")
			content = ""
			continue
		}

		content += line + "\n"
	}
	if header != "" {
		sections[header] = content // Push the final section, if there is one
	}

	for header, content := range sections {
		if strings.Contains(strings.ToLower(header), "spec") {
			tc.SpecMd = content
		} else if strings.Contains(strings.ToLower(header), "steps") {
			tc.StepsMd = content
		} else if strings.Contains(strings.ToLower(header), "expected result") {
			tc.ExpectedResultMd = content
		}
	}
	return tc, nil
}
