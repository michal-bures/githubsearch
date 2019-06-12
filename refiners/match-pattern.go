package refiners

import (
	"github.com/google/go-github/github"
	"github.com/minio/minio/pkg/wildcard"
)

type MatchPattern struct {
	Pattern string
}

// removes all results that don't contain exact match with the pattern
func (m MatchPattern) Apply(results *[]github.CodeResult) *[]github.CodeResult {
	return filterResultsWith(results, getPatternFilterFunc(m.Pattern))
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
