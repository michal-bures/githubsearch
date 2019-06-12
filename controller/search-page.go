package controller

import (
	"html/template"
	"net/http"
)

type SearchResult struct {
	Repository *string
	Name       *string
	Path       *string
	FileUrl    *string
	Score      int
	Fragments  []*string
}

type SearchPageData struct {
	ShowResults    bool
	SearchString   string
	SearchLanguage string
	Results        []SearchResult
}

var searchPageTemplate = template.Must(template.ParseFiles("templates/search-page.html"))

func SearchPage(w http.ResponseWriter, params SearchPageData) {
	err := searchPageTemplate.Execute(w, params)
	handleError(w, err)
}
