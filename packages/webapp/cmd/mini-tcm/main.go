package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/bensivo/mini-tcm/packages/webapp/pkg/testcase"
	"github.com/spf13/pflag"
)

func main() {
	var dir = pflag.StringP("dir", "d", ".", "The directory to load test cases from")
	pflag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to load template index.html", http.StatusInternalServerError)
			return
		}

		tcs, err := testcase.LoadFromDir(*dir)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "failed to parse test cases", http.StatusInternalServerError)
		}

		type TestCaseRendered struct {
			Id                 string
			Name               string
			SpecMd             string
			SpecHtml           template.HTML
			StepsMd            string
			StepsHtml          template.HTML
			ExpectedResultMd   string
			ExpectedResultHtml template.HTML
		}

		tcsRendered := make([]TestCaseRendered, len(tcs))
		for i := range tcs {
			// TODO: Find image references, and find a way to serve those files as well

			// TODO: Use a HTML sanitizer like bluemonday
			tcsRendered[i] = TestCaseRendered{
				Id:                 tcs[i].Id,
				Name:               tcs[i].Name,
				SpecMd:             tcs[i].SpecMd,
				SpecHtml:           template.HTML(mdToHTML([]byte(tcs[i].SpecMd))),
				StepsMd:            tcs[i].StepsMd,
				StepsHtml:          template.HTML(mdToHTML([]byte(tcs[i].StepsMd))),
				ExpectedResultMd:   tcs[i].ExpectedResultMd,
				ExpectedResultHtml: template.HTML(mdToHTML([]byte(tcs[i].ExpectedResultMd))),
			}
		}

		t.Execute(w, struct {
			TestCases []TestCaseRendered
		}{
			TestCases: tcsRendered,
		})
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
