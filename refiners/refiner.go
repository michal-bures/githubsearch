package refiners

import (
	"context"
	"github.com/google/go-github/github"
)

type SearchResultsRefiner interface {
	Apply(ctx context.Context, results *[]github.CodeResult) *[]github.CodeResult //TODO why do you actually pass a context here? I haven't seen it used
}
