// home.go runs the homepage,
// provides the user form,
// stores the response in
// in some format, crawls,
// then graphs the crawl.
//TODO add cookies and past starting urls
package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Crawl struct{
	Url string			// Start
	Keyword string	// Optional
	Type string			// "B" or "D"
	BL string				// Breadth limit
	DL string				// Depth limit
}

type Graph struct{
	//
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	crawl := Crawl{
		Url: r.FormValue("Url"),
		Keyword: r.FormValue("Keyword"),
		Type: r.FormValue("Type"),
		BL: r.FormValue("BL"),
		DL: r.FormValue("DL"),
	}
	// crawl is now populated.
	fmt.Printf("%+v\n", crawl) // debug

	// TODO format crawl settings as needed
	_ = crawl

	// TODO
	/* Crawler Program
			input: crawl (formatted)
			output: graph (formatted) */

	// TODO
	/* Render graph
		input: graph (formatted)
		output: D3.js, graphs.html? */

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
