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

// Stores parsed html templates, so we don't have to parse them every time we render a page
var tmpl *template.Template

func (s *Server) Serve() {
	// Parse all templates in the templates directory, save in gloval variable
	t, err := template.New("").ParseGlob("templates/*.template.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl = t

	// Define routes, and start server
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HomepageHandler) // TODO: change the homepage to just a splash screen with 2 options "Manage Test Cases", "Run Test Cases"
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// TODO: add another route /testcases/{path+}, which serves all the test cases in taht path, and keeps everything relative to that path

	host := fmt.Sprintf(":%d", s.Port)
	fmt.Println("Listening on", host)
	http.ListenAndServe(host, mux)

}

func (s *Server) HomepageHandler(w http.ResponseWriter, r *http.Request) {
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

	err = tmpl.ExecuteTemplate(w, "homepage", struct {
		TestCases []TestCaseRendered
		Folder    *folders.Folder
	}{
		TestCases: tcsRendered,
		Folder:    folder,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
		return
	}
}
