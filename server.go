package main

import (
	_ "context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "artus/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMonitorServiceServer
}

func (s *server) Monitor(stream pb.MonitorService_MonitorServer) error {
	log.Println("New Monitor stream connected")
	// Start a goroutine to send periodic ServerMessage to client (optional)
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			err := stream.Send(&pb.ServerMessage{
				Text:    "server heartbeat",
				Command: "",
			})
			if err != nil {
				log.Println("Error sending heartbeat:", err)
				return
			}
		}
	}()

	// receive metrics from client
	for {
		metrics, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed stream")
			return nil
		}
		if err != nil {
			log.Println("stream.Recv error:", err)
			return err
		}

		// print summary
		ts := metrics.GetTs()
		tStr := "<nil>"
		if ts != nil {
			tStr = ts.AsTime().Format(time.RFC3339)
		}
		log.Printf("Recv from %s @ %s: CPU cores=%d, cpu[0] usage=%.2f%%, mem used=%.2f%%, disk used=%.2f%%\n",
			metrics.GetHostname(),
			tStr,
			len(metrics.GetCpus()),
			func() float64 {
				if len(metrics.GetCpus()) > 0 {
					return metrics.GetCpus()[0].GetUsagePercent()
				}
				return 0
			}(),
			metrics.GetMemory().GetUsedPercent(),
			metrics.GetDisk().GetUsedPercent(),
		)

		// Optionally respond with an ack
		err = stream.Send(&pb.ServerMessage{
			Text:    fmt.Sprintf("ack: received metrics at %s", time.Now().Format(time.RFC3339)),
			Command: "",
		})
		if err != nil {
			log.Println("Error sending ack:", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMonitorServiceServer(s, &server{})
	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
