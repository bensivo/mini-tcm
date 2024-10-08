package testcase_test

import (
	"errors"
	"testing"

	"github.com/bensivo/mini-tcm/packages/webapp/pkg/testcase"
)

var FromMarkdownTestData = []struct {
	name string
	md   string
	tc   testcase.TestCase
	err  error
}{
	{
		name: "Happy Path",
		md: `
---
id: TC-1
name: test-name
---

## Spec
here is my spec

## Steps
here are my steps

## Expected Result
here are my expected results
`,
		tc: testcase.TestCase{
			Id:               "TC-1",
			Name:             "test-name",
			SpecMd:           "here is my spec",
			StepsMd:          "here are my steps",
			ExpectedResultMd: "here are my expected results",
		},
		err: nil,
	},
	{
		name: "Missing Front-Matter",
		md: `
## Spec
here is my spec

## Steps
here are my steps
		`,
		tc:  testcase.TestCase{},
		err: testcase.ErrNoYamlFrontMatter,
	},
	{
		name: "Missing sections",
		md: `
---
id: TC-1
name: test-name
---
## Spec
here is my spec
	`,
		tc: testcase.TestCase{
			Id:               "TC-1",
			Name:             "test-name",
			SpecMd:           "here is my spec",
			StepsMd:          "",
			ExpectedResultMd: "",
		},
		err: nil,
	},
	{
		name: "Just frontmatter",
		md: `
---
id: TC-1
name: test-name
---
`,
		tc: testcase.TestCase{
			Id:               "TC-1",
			Name:             "test-name",
			SpecMd:           "",
			StepsMd:          "",
			ExpectedResultMd: "",
		},
		err: nil,
	},
}

func TestFromMarkdown(t *testing.T) {
	for _, testdata := range FromMarkdownTestData {
		t.Run(testdata.name, func(t *testing.T) {
			tc, err := testcase.FromMarkdown(testdata.md)
			if !errors.Is(err, testdata.err) {
				t.Errorf("Expected error to be '%v', got '%v'", testdata.err, err)
			}

			if tc.Id != testdata.tc.Id {
				t.Errorf("Expected Id to be '%s', got %s", testdata.tc.Id, tc.Id)
			}

			if tc.Name != testdata.tc.Name {
				t.Errorf("Expected Name to be '%s', got %s", testdata.tc.Name, tc.Name)
			}

			if tc.SpecMd != testdata.tc.SpecMd {
				t.Errorf("Expected SpecMd to be '%s', got '%s'", testdata.tc.SpecMd, tc.SpecMd)
			}

			if tc.StepsMd != testdata.tc.StepsMd {
				t.Errorf("Expected StepsMd to be '%s', got %s", testdata.tc.StepsMd, tc.StepsMd)
			}
		})
	}

}
