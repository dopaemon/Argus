package server

import (
	"fmt"
	"log"
	"net"

	"github.com/dopaemon/artus/internal/gRPC/metrics"
	"google.golang.org/grpc"
)

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	metrics.RegisterMetricsServiceServer(grpcServer, &MetricsServer{})

	fmt.Println("gRPC server listening on", port)
	go grpcServer.Serve(lis)
}
