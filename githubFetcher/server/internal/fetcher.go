package internal

import (
	"context"
	"encoding/json"
	"fmt"
	proto2 "rigSecurityMaytar/githubFetcher/proto"

	"rigSecurityMaytar/utils/github"
)

type GithubFetcher struct {
	proto2.UnimplementedGithubFetcherServer
	GithubClient *github.Client
}

// GetRepositoriesEvaluation fetches all repositories and evaluate their configurations
func (f *GithubFetcher) GetRepositoriesEvaluation(ctx context.Context, req *proto2.GetRepositoriesEvaluationRequest) (*proto2.GetRepositoriesEvaluationResponse, error) {
	ctx = context.WithValue(ctx, "organization", req.Organization)

	repos, err := f.getRepos(ctx, req.Organization)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to get repositories: %w", ctx, err)
	}

	reposResults := make([]*proto2.Repository, 0)

	for _, repo := range repos {
		regoInput, err := structToRegoInput(ctx, &repo)
		if err != nil {
			return nil, fmt.Errorf("[%s] failed to convert rego input for repo %s: %w", ctx, repo.Name, err)
		}

		policiesResult, err := evaluatePolicies(ctx, regoInput)
		if err != nil {
			return nil, fmt.Errorf("[%s] failed to run the policies on the repo %s: %w", ctx, repo.Name, err)
		}

		repoResult := &proto2.Repository{
			Name:     repo.Name,
			Policies: policiesResult,
		}

		reposResults = append(reposResults, repoResult)
	}

	return &proto2.GetRepositoriesEvaluationResponse{Repositories: reposResults}, nil
}

func (f *GithubFetcher) getRepos(ctx context.Context, organization string) ([]Repo, error) {
	rawRepos, err := f.fetchRepos(ctx, organization)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch repositories names for the organization %s: %w", ctx, organization, err)
	}

	repos := make([]Repo, 0)

	for _, rawRepo := range rawRepos {
		repoName := rawRepo.Name

		rawCollaborators, err := f.fetchRepoCollaborators(ctx, organization, repoName)
		if err != nil {
			return nil, fmt.Errorf("[%s] failed to fetch repo collaborators: %w", ctx, err)
		}

		repoCollaborators := Repo{
			Name:          repoName,
			Collaborators: rawCollaborators,
		}

		repos = append(repos, repoCollaborators)
	}

	return repos, nil
}

func (f *GithubFetcher) fetchRepos(ctx context.Context, organization string) ([]RawRepo, error) {
	resp, err := f.GithubClient.GetRepositories(ctx, organization)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch repos names for organization [%s]: %w", ctx, organization, err)
	}

	var repos []RawRepo

	err = json.Unmarshal(resp.Body(), &repos)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to parse response: %w", ctx, err)
	}

	return repos, nil
}

func (f *GithubFetcher) fetchRepoCollaborators(ctx context.Context, organization string, repoName string) ([]RawCollaborator, error) {
	ctx = context.WithValue(ctx, "repoName", repoName)

	resp, err := f.GithubClient.GetCollaborators(ctx, organization, repoName)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch repo collaborators for repo %s: %w", ctx, repoName, err)
	}

	var collaborators []RawCollaborator

	err = json.Unmarshal(resp.Body(), &collaborators)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to parse collaborators response: %w", ctx, err)
	}

	return collaborators, nil
}

func structToRegoInput(ctx context.Context, repo interface{}) (map[string]interface{}, error) {
	marshaledRepo, err := json.Marshal(repo)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to marshal struct to JSON: %w", ctx, err)
	}

	var regoInput map[string]interface{}

	if err := json.Unmarshal(marshaledRepo, &regoInput); err != nil {
		return nil, fmt.Errorf("[%s] failed to unmarshal: %w", ctx, err)
	}

	return regoInput, nil
}
