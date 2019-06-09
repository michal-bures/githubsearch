package controller

import (
	"fmt"
	"github-search/searcher"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"html/template"
	"log"
	"net/http"
)

type Controller struct {
	searchEngine searcher.Searcher
}

var resultsTemplate = template.Must(template.ParseFiles("templates/results.html"))
var searchPageTemplate = template.Must(template.ParseFiles("templates/search-form.html"))

func (c *Controller) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	keywords := r.URL.Query()["search"]
	if keywords == nil {
		SearchFormPage(w)
	} else {
		results, err := c.searchEngine.Search(keywords[0])
		handleError(w, err)

		ResultsPage(w, ResultsPageData{
			SearchString: r.URL.Query()["search"][0],
			Results:      resultsToString(results),
		})
	}
}

func handleError(w http.ResponseWriter, e error) {
	if e != nil {
		log.Fatal(errors.WithStack(e))
		w.WriteHeader(500)
		fmt.Fprint(w, "Oops, something went wrong")
	}
}

func resultsToString(codeResults []github.CodeResult) []string {
	searchResults := make([]string, len(codeResults))

	fmt.Printf("Total results: %d", len(codeResults))
	fmt.Printf("%+v\n", codeResults[0])

	for i, codeResult := range codeResults {
		searchResults[i] = *codeResult.Repository.Name
	}
	return searchResults
}

func NewController(searchEngine searcher.Searcher) *Controller {
	return &Controller{
		searchEngine: searchEngine,
	}
}
