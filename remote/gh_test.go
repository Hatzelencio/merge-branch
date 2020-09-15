package remote

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-github/v32/github"
	"github.com/hatzelencio/merge-branch/utils/mocks"
	"golang.org/x/net/context"
)

type envVariables struct {
	token         string
	repository    string
	head          string
	base          string
	commitMessage string
}

func init() {
	cli = NewGithubClient(nil, &mocks.MockClient{})
}

func TestMergeSuccess(t *testing.T) {
	params := envVariables{
		token:         "secret-token",
		repository:    "owner/owner-repo",
		head:          "aae1a1fef...",
		base:          "heads/branch,heads/branch-two",
		commitMessage: "Merged aae1a1fef into head",
	}

	setEnvVariables(params)
	mockMergeSuccess()

	err := Merge()
	if err != nil {
		t.Fatal("it can'n be merged.")
	}
}

func TestMergeNoContent(t *testing.T) {
	params := envVariables{
		token:         "secret-token",
		repository:    "owner/owner-repo",
		head:          "aae1a1fef...",
		base:          "heads/branch,heads/branch-two",
		commitMessage: "Merged aae1a1fef into head",
	}

	setEnvVariables(params)
	mockMergeNoContent()

	err := Merge()
	if err != nil {
		t.Fatal("it can'n be merged.")
	}
}

func TestMergeConflict(t *testing.T) {
	params := envVariables{
		token:         "secret-token",
		repository:    "owner/owner-repo",
		head:          "aae1a1fef...",
		base:          "heads/branch,heads/branch-two",
		commitMessage: "Merged aae1a1fef into head",
	}

	setEnvVariables(params)
	mockMergeConflict()

	err := Merge()
	if err == nil || err.Error() != "merge conflict" {
		t.Fatal(err)
	}
}

func TestMergeBaseNotFound(t *testing.T) {
	params := envVariables{
		token:         "secret-token",
		repository:    "owner/owner-repo",
		head:          "aae1a1fef...",
		base:          "heads/branch,heads/branch-two",
		commitMessage: "Merged aae1a1fef into head",
	}

	setEnvVariables(params)
	mockMergeBaseNotFound()

	err := Merge()
	if err == nil || err.Error() != "base does not exist" {
		t.Fatal(err)
	}
}

func TestMergeHeadNotFound(t *testing.T) {
	params := envVariables{
		token:         "secret-token",
		repository:    "owner/owner-repo",
		head:          "aae1a1fef...",
		base:          "heads/branch,heads/branch-two",
		commitMessage: "Merged aae1a1fef into head",
	}

	setEnvVariables(params)
	mockMergeHeadNotFound()

	err := Merge()
	if err == nil || err.Error() != "head does not exist" {
		t.Fatal(err)
	}
}

func setEnvVariables(env envVariables) {
	_ = os.Setenv("GITHUB_TOKEN", env.token)
	_ = os.Setenv("GITHUB_REPOSITORY", env.repository)
	_ = os.Setenv("GITHUB_REF", env.head)
	_ = os.Setenv("INPUT_BASE", env.base)
	_ = os.Setenv("INPUT_COMMIT_MESSAGE", env.commitMessage)
}

func mockMergeSuccess() {
	mocks.PostMergeFunc = func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
		res := github.Response{
			Response: &http.Response{StatusCode: 201},
		}
		return &github.RepositoryCommit{}, &res, nil
	}
}

func mockMergeNoContent() {
	mocks.PostMergeFunc = func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
		res := github.Response{
			Response: &http.Response{StatusCode: 204},
		}
		return &github.RepositoryCommit{}, &res, nil
	}
}

func mockMergeConflict() {
	mocks.PostMergeFunc = func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
		res := github.Response{
			Response: &http.Response{StatusCode: 409},
		}
		return nil, &res, fmt.Errorf("merge conflict")
	}
}

func mockMergeBaseNotFound() {
	mocks.PostMergeFunc = func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
		res := github.Response{
			Response: &http.Response{StatusCode: 404},
		}
		return nil, &res, fmt.Errorf("base does not exist")
	}
}

func mockMergeHeadNotFound() {
	mocks.PostMergeFunc = func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
		res := github.Response{
			Response: &http.Response{StatusCode: 404},
		}
		return nil, &res, fmt.Errorf("head does not exist")
	}
}
