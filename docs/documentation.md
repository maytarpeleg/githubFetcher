# GitHub Repositories Evaluator

## Overview
The Github Repositories Evaluator is a gRPC-based backend service developed in Golang to validate policies for GitHub organization repositories.

The service is designed to ensure that specific access rules, such as prohibiting certain usernames from having repository access, are enforced using Open Policy Agent (OPA).

This service fetches repository details from GitHub, evaluates the repositories against a set of policies stored in the policies directory, and returns evaluation results.

A simple client is provided to interact with the service.


## GitHub API Endpoints
The `githubFetcher` service uses these GitHub endpoints.
- `orgs/<org>/repos` - fetches the repositories of a specific organization.
- `repos/<org>/<repo>/collaborators` - fetches the collaborators of a specific repository.

### Testing
Install and use **Postman** for testing the endpoints.

## Regenerating Protocol Buffers
`protoc --go_out=. --go-grpc_out=. proto/server.proto`


## Future Enhancements
### Pagination Support
Implement handling for paginated responses from the GitHub API to efficiently fetch large datasets.
### Improved Error Handling
Add robust error-handling mechanisms to address scenarios like:
- API rate limits
- Missing or invalid access tokens
- Network failures
### Additional Rego Policies
Add Rego policies to strengthen the security and compliance checks.
