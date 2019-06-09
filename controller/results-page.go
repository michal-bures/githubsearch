package controller

import "net/http"

type ResultsPageData struct {
	SearchString string
	Results      []string
}

func ResultsPage(w http.ResponseWriter, params ResultsPageData) {
	err := resultsTemplate.Execute(w, params)
	handleError(w, err)
}
