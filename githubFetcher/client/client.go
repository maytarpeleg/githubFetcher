package client

import (
	"context"
	"fmt"
	"log"
	"rigSecurityMaytar/githubFetcher/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGRPCClient(ctx context.Context, address string) (proto.GithubFetcherClient, func(), error) {
	ctx = context.WithValue(ctx, "address", address)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s] could not create Github fetcher client: %w", ctx, err)
	}
	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("fail to close connection: %v", err)
		}
	}

	return proto.NewGithubFetcherClient(conn), cleanup, nil
}
