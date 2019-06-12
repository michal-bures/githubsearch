package refiners

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"sort"
)

type SortByRepositoryScore struct {
	MaxRequests int
	Client      *github.Client
}

//sorts results based on repository score
func (s SortByRepositoryScore) Apply(results *[]github.CodeResult) *[]github.CodeResult {
	scores := make([]int, s.MaxRequests)

	for i := 0; i < smallerOf(s.MaxRequests, len(*results)); i++ {
		scores[i] = s.getRepositoryScore((*results)[i].Repository)
	}

	var sorter = ByScoreFromHighest{
		Results: *results,
		Scores:  scores,
	}

	sort.Sort(sorter)

	return &sorter.Results
}

func (s SortByRepositoryScore) getRepositoryScore(repo *github.Repository) int {

	owner := (*repo).Owner.GetLogin()
	name := (*repo).GetName()

	fmt.Printf("Owner %s, Name %s\n", owner, name)

	repository, _, err := s.Client.Repositories.Get(context.Background(), owner, name)

	if err != nil {
		panic(err)
	}

	log.Print("GOT REPO")
	log.Println(repository)

	score := (*repository).GetWatchersCount()

	fmt.Printf("Repo %s has score %d\n", (*repo).Name, score)

	return score
}

func smallerOf(int1 int, int2 int) int {
	if int1 <= int2 {
		return int1
	}
	return int2
}

type ByScoreFromHighest struct {
	Results []github.CodeResult
	Scores  []int
}

func (a ByScoreFromHighest) Len() int { return len(a.Results) }
func (a ByScoreFromHighest) Swap(i, j int) {
	a.Results[i], a.Results[j] = a.Results[j], a.Results[i]
	a.Scores[i], a.Scores[j] = a.Scores[j], a.Scores[i]
}
func (a ByScoreFromHighest) Less(i, j int) bool {
	return a.Scores[i] < a.Scores[j]
}
