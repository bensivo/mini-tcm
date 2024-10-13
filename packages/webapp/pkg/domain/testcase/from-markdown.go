package testcase

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrNoYamlFrontMatter = fmt.Errorf("no yaml front matter found")
)

func FromMarkdown(md string) (TestCase, error) {
	tc := TestCase{}

	// Extract the yaml front matter
	headingRegex := regexp.MustCompile(`(?s)---\n(.*?)\n---\n`)
	match := headingRegex.FindStringSubmatch(md)
	if len(match) < 2 {
		fmt.Println("no yaml front matter found")
		return tc, ErrNoYamlFrontMatter
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

	sections := map[string]string{} // Maps section headers to the content of that section
	var header string
	var content string
	lines := strings.Split(md, "\n")

	headingRegex = regexp.MustCompile(`^#+\s`) // Regex to check for md headers
	for _, line := range lines {
		match := headingRegex.MatchString(line) // See if we hit a new heading
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
		content = strings.TrimSpace(content)
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
