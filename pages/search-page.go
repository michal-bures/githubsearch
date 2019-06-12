package pages

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

var searchPageTemplate = template.Must(template.ParseFiles("pages/search-page.html"))

func SearchPage(w http.ResponseWriter, errorHandler func(e error), params SearchPageData) {
	err := searchPageTemplate.Execute(w, params)
	errorHandler(err)
}
