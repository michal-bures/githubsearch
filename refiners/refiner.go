package refiners

import "github.com/google/go-github/github"

type SearchResultsRefiner interface {
	Apply(results *[]github.CodeResult) *[]github.CodeResult
}

type filterFunction = func(result *github.CodeResult) bool

func filterResultsWith(results *[]github.CodeResult, filterFunction filterFunction) *[]github.CodeResult {
	filteredResults := make([]github.CodeResult, 0, len(*results))
	for _, candidate := range *results {
		if filterFunction(&candidate) {
			filteredResults = append(filteredResults, candidate)
		}
	}
	return &filteredResults
}
