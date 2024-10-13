package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bensivo/mini-tcm/packages/webapp/pkg/domain/folders"
	"github.com/bensivo/mini-tcm/packages/webapp/pkg/domain/testcase"
)

type Server struct {
	Port        int
	TestCaseDir string
}

func (s *Server) Serve() {
	mux := http.NewServeMux()

	// TODO: change the homepage to just a splash screen with 2 options "Manage Test Cases", "Run Test Cases"
	mux.HandleFunc("/", s.HomepageHandler)

	// TODO: add another route /testcases/{path+}, which serves all the test cases in taht path, and keeps everything relative to that path

	host := fmt.Sprintf(":%d", s.Port)
	http.ListenAndServe(host, mux)

}

func (s *Server) HomepageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.template.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to load template index.html", http.StatusInternalServerError)
		return
	}

	tcs, err := testcase.LoadFromDir(s.TestCaseDir)
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

	folder, err := folders.LoadFromFs(s.TestCaseDir)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to load folders", http.StatusInternalServerError)
		return
	}

	t.Execute(w, struct {
		TestCases []TestCaseRendered
		Folder    *folders.Folder
	}{
		TestCases: tcsRendered,
		Folder:    folder,
	})
}
