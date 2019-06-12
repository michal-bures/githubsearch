package refiners

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/minio/minio/pkg/wildcard"
)

type MatchPattern struct {
	Pattern string
}

// removes all results that don't contain exact match with the pattern
func (m MatchPattern) Apply(ctx context.Context, results *[]github.CodeResult) *[]github.CodeResult {
	return filterResultsWith(results, getPatternFilterFunc(m.Pattern))
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

func getPatternFilterFunc(pattern string) func(result *github.CodeResult) bool {
	return func(result *github.CodeResult) bool {
		for _, match := range result.TextMatches {
			if wildcard.Match("*"+pattern+"*", *match.Fragment) {
				return true
			}
		}
		return false
	}
}
