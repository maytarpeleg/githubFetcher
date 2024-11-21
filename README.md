# GithubFetcher

A gRPC-based backend service for validating policies on GitHub organization repositories.

The service provides one RPC method - GetRepositoriesEvaluation, which receives a GitHub organization name, fetches its repositories details and evaluates them.
Evaluation is done by the OPA engine. Policies are read from the `policies` directory.

A simple client library is also provided.

# Prerequisites
* Go 1.18+ 
* protoc (Protocol Buffers Compiler)
* protoc go plugins (See instructions in https://grpc.io/docs/languages/go/quickstart/)

# Install
* `go mod tidy`

# Usage
* Create a GitHub classic API token 
  * Follow instruction in the following link and check the repo scope access https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens
* Make sure you have the right permissions to the wanted Github organization
* `GITHUB_TOKEN=<token> go run githubFetcher/server/cmd/main.go`

# Policies
### Current Policies
1. GitHub repository collaborators should not be in the blocked list
2. GitHub repository collaborators site-admin should be in the allowed list

