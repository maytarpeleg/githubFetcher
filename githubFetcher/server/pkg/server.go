package pkg

import (
	"context"
	"fmt"
	"log"
	"net"
	"rigSecurityMaytar/githubFetcher/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"rigSecurityMaytar/githubFetcher/server/internal"
	"rigSecurityMaytar/utils/github"
)

func StartGRPCServer(ctx context.Context, address string) error {
	ctx = context.WithValue(ctx, "address", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("[%s] failed to listen on %s: %w", ctx, address, err)
	}

	grpcServer := grpc.NewServer()

	githubClient, err := github.NewClient()
	if err != nil {
		return fmt.Errorf("[%s] failed to create utils client: %w", ctx, err)
	}

	proto.RegisterGithubFetcherServer(grpcServer, &internal.GithubFetcher{GithubClient: githubClient})

	reflection.Register(grpcServer)

	log.Printf("Server listening on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("[%s] failed to start gRPC server: %w", ctx, err)
	}

	return nil
}
