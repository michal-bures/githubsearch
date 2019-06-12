package searcher

import (
	"github.com/google/go-github/github"
)

type Searcher interface {
	Search(query string, language string) (*[]github.CodeResult, error)
}
