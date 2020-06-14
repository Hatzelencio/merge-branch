package mocks

import (
	"github.com/google/go-github/v32/github"
	"golang.org/x/net/context"
)

// MockClient mock of github.Client
type MockClient struct {
	MergeFunc func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error)
}

var (
	// PostMergeFunc mock
	PostMergeFunc func(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error)
)


// Merge mock
func (m *MockClient) Merge(ctx context.Context, owner, repo string, request *github.RepositoryMergeRequest) (*github.RepositoryCommit, *github.Response, error) {
	return PostMergeFunc(ctx, owner, repo, request)
}