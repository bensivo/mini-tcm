package main

import (
	"fmt"
	"html/template"
	"net/http"

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

		t.Execute(w, struct {
			TestCases []testcase.TestCase
		}{
			TestCases: tcs,
		})
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
