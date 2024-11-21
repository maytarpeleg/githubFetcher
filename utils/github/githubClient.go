package github

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
	token  string
}

func NewClient() (*Client, error) {
	const githubTokenEnvVarName = "GITHUB_TOKEN"

	token := os.Getenv(githubTokenEnvVarName)
	if token == "" {
		return nil, fmt.Errorf("%s environment variable must be set", githubTokenEnvVarName)
	}

	client := &Client{
		client: resty.New(),
		token:  token,
	}

	return client, nil
}

func (c *Client) getURL(ctx context.Context, suffixURL string) (*resty.Response, error) {
	const baseUrl = "https://api.github.com"

	url := fmt.Sprintf("%s/%s?per_page=100", baseUrl, suffixURL)

	ctx = context.WithValue(ctx, "url", url)

	resp, err := c.client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.token)).
		SetHeader("Accept", "application/vnd.github+json").
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch from Github: %w", ctx, err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("[%s] failed to fetch from Github: status code %s", ctx, resp.Status())
	}

	return resp, nil
}

func (c *Client) GetRepositories(ctx context.Context, organization string) (*resty.Response, error) {
	suffixURL := fmt.Sprintf("orgs/%s/repos", organization)

	resp, err := c.getURL(ctx, suffixURL)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch repos names for organization [%s]: %w", ctx, organization, err)
	}

	return resp, nil
}

func (c *Client) GetCollaborators(ctx context.Context, organization string, repoName string) (*resty.Response, error) {
	suffixURL := fmt.Sprintf("repos/%s/%s/collaborators", organization, repoName)

	resp, err := c.getURL(ctx, suffixURL)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to fetch collaborators for the repo [%s]: %w", ctx, repoName, err)
	}

	return resp, nil
}
