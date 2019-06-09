package searcher

import (
	"github.com/google/go-github/github"
)

type Searcher interface {
	Search(keywords string) ([]github.CodeResult, error)
}
