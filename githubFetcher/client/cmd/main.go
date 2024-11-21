package main

import (
	"context"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
	"log"
	"os"
	githubFetcherClient "rigSecurityMaytar/githubFetcher/client"
	"rigSecurityMaytar/githubFetcher/proto"
)

const (
	addressEnvVarName = "ADDRESS"
	defaultAddress    = "localhost:9900"
)

func main() {
	ctx := context.TODO()
	var args struct {
		Organization string `arg:"required"`
	}
	arg.MustParse(&args)

	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load .env file: %v", err)
	}

	address := os.Getenv(addressEnvVarName)
	if address == "" {
		log.Printf("%s environment variable was not set, using default %s", addressEnvVarName, defaultAddress)
		address = defaultAddress
	}

	client, cleanup, err := githubFetcherClient.GetGRPCClient(ctx, address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer cleanup()

	req := &proto.GetRepositoriesEvaluationRequest{
		Organization: args.Organization,
	}

	resp, err := client.GetRepositoriesEvaluation(ctx, req)
	if err != nil {
		log.Fatalf("Failed to fetch repository evaluations: %v", err)
	}

	if len(resp.Repositories) == 0 {
		fmt.Printf("Expected repositories, got none")
		return
	}

	for _, repo := range resp.Repositories {
		log.Printf("Repo: %s", repo.Name)
		for _, policy := range repo.Policies {
			log.Printf("Policy ID: %s, Result: %t, Title: %s", policy.Id, policy.Result, policy.Title)
		}
	}
}
