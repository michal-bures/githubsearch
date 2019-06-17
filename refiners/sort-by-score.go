package refiners

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"sort"
)

type SortByRepositoryScore struct {
	MaxRequests int
	Client      *github.Client
}

//sorts results based on repository score
func (s SortByRepositoryScore) Apply(ctx context.Context, results *[]github.CodeResult) *[]github.CodeResult {
	numberOfScoredResults := smaller(s.MaxRequests, len(*results))

	scores := make([]int, numberOfScoredResults)

	for i := 0; i < numberOfScoredResults; i++ {
		scores[i] = s.getRepositoryScore(ctx, (*results)[i].Repository)
	}

	var sorter = ByScoreFromHighest{
		Results: *results,
		Scores:  scores,
	}

	sort.Sort(sorter)

	return &sorter.Results
}

func (s SortByRepositoryScore) getRepositoryScore(ctx context.Context, repo *github.Repository) int {
	owner := (*repo).Owner.GetLogin()
	name := (*repo).GetName()

	repository, _, err := s.Client.Repositories.Get(context.Background(), owner, name)

	if err != nil {
		fmt.Printf("Failed to retrieve repository metadata for %s/%s", owner, name)
		return 0
	}

	score := repository.GetWatchersCount() //TODO get actually already returns a pointer, so there's no need for it here

	return score
}

func smaller(int1 int, int2 int) int {
	if int1 <= int2 {
		return int1
	}
	return int2
}

type ByScoreFromHighest struct {
	Results []github.CodeResult
	Scores  []int
}

func (a ByScoreFromHighest) Len() int {
	return len(a.Scores)
}
func (a ByScoreFromHighest) Swap(i, j int) {
	a.Results[i], a.Results[j] = a.Results[j], a.Results[i]
	a.Scores[i], a.Scores[j] = a.Scores[j], a.Scores[i]
}
func (a ByScoreFromHighest) Less(i, j int) bool {
	return a.Scores[i] > a.Scores[j]
}
