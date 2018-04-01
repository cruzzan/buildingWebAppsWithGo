package main

import (
	"net/http"
	"github.com/russross/blackfriday"
)

func main()  {
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8181", nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("raw-markdown")))
	rw.Write(markdown)
}
