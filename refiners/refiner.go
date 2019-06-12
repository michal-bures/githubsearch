package refiners

import (
	"context"
	"github.com/google/go-github/github"
)

type SearchResultsRefiner interface {
	Apply(ctx context.Context, results *[]github.CodeResult) *[]github.CodeResult
}
