package main

import (
	"log"
	"time"

	clientMetrics "github.com/dopaemon/artus/internal/metrics"
	grpcMetrics "github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc"
)

func main() {
	apiKey := "889a8c37375f977c3a8f7428a1bab518435ac834f210fa64182f2768105a7eb5"

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpcMetrics.NewMetricsServiceClient(conn)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		clientMetrics.Send(client, apiKey)
	}
}
