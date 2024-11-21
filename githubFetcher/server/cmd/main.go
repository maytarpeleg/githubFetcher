package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	server "rigSecurityMaytar/githubFetcher/server/pkg"
)

const (
	addressEnvVarName = "ADDRESS"
	defaultAddress    = "localhost:9900"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load .env file: %v", err)
	}

	address := os.Getenv(addressEnvVarName)
	if address == "" {
		log.Printf("%s environment variable was not set, using default %s", addressEnvVarName, defaultAddress)
		address = defaultAddress
	}

	if err := server.StartGRPCServer(context.TODO(), address); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
