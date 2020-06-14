package remote

import (
	"fmt"
	"github.com/google/go-github/v32/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
)

const (
	token  = "GITHUB_TOKEN"
	ghref  = "GITHUB_REF"
	ghrepo = "GITHUB_REPOSITORY"
	base   = "INPUT_BASE"
	head   = "INPUT_HEAD"
)

var (
	ctx context.Context
	cli GithubClient
)

func init() {
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(token)},
	)
	tc := oauth2.NewClient(ctx, ts)
	cli = NewGithubClient(tc, nil)
}

// GithubGitService interface
type GithubGitService interface {
	Merge(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error)
}

// GithubClient is a wrapper of github.Client
type GithubClient struct {
	Repositories GithubGitService
	*github.Client
}

// NewGithubClient Create a new github client
func NewGithubClient(client *http.Client, repoMock GithubGitService) GithubClient {
	if repoMock != nil {
		return GithubClient{
			Repositories: repoMock,
		}
	}

	cli := github.NewClient(client)
	return GithubClient{
		Repositories: cli.Repositories,
	}
}

// ValidateInputs validate if GITHUB_TOKEN and INPUT_REFS are present like environment variables
func ValidateInputs() error {
	if len(os.Getenv(token)) == 0 {
		return fmt.Errorf("%v is required env variable to trigger this action", token)
	}

	if len(os.Getenv(base)) == 0 {
		return fmt.Errorf("%v is required input to trigger this action", "base")
	}

	return nil
}

func getRefBase() string {
	return strings.Replace(os.Getenv(base), " ", "", -1)
}

func getRefHead() string {
	var ref string

	if len(os.Getenv(head)) == 0 {
		ref = strings.Replace(os.Getenv(ghref), " ", "", -1)
	} else {
		ref = strings.Replace(os.Getenv(head), " ", "", -1)
	}

	return ref
}

func getOwnerRepo() (string, string) {
	ownerRepo := strings.Split(os.Getenv(ghrepo), "/")
	return ownerRepo[0], ownerRepo[1]
}

// Merge create a merge on remote
func Merge() error {
	refBase := getRefBase()
	refHead := getRefHead()
	owner, repo := getOwnerRepo()

	req := github.RepositoryMergeRequest{
		Base: &refBase,
		Head: &refHead,
	}
	_, res, err := cli.Repositories.Merge(ctx, owner, repo, &req)

	if res == nil {
		return err
	}

	if res.StatusCode == http.StatusConflict || res.StatusCode == http.StatusNotFound {
		return err
	}

	return nil
}
