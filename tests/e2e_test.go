package tests

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	githubFetcherServer "rigSecurityMaytar/githubFetcher/server/pkg"
	"testing"

	githubFetcherClient "rigSecurityMaytar/githubFetcher/client"
	"rigSecurityMaytar/githubFetcher/proto"
)

const (
	testGithubOrg = "testorgmaytar"
	address       = "localhost:9900"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load .env file: %v", err)
	}

	go func() {
		if err := githubFetcherServer.StartGRPCServer(context.Background(), address); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
}

func TestE2E_GetReposEvaluations(t *testing.T) {
	ctx := context.TODO()

	client, cleanup, err := githubFetcherClient.GetGRPCClient(ctx, address)
	if err != nil {
		t.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer cleanup()

	req := &proto.GetRepositoriesEvaluationRequest{
		Organization: testGithubOrg,
	}

	resp, err := client.GetRepositoriesEvaluation(ctx, req)
	if err != nil {
		t.Fatalf("Failed to fetch repository evaluations: %v", err)
	}

	if len(resp.Repositories) == 0 {
		t.Errorf("Expected repositories, got none")
	}

	for _, repo := range resp.Repositories {
		t.Logf("Repo: %s", repo.Name)
		for _, policy := range repo.Policies {
			t.Logf("Policy ID: %s, Result: %t, Title: %s", policy.Id, policy.Result, policy.Title)
		}
	}
}
