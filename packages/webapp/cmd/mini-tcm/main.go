package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	t, err := template.ParseFiles("templates/index.html")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	t.Execute(w, nil)
	// })

	mux.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", mux)
}
