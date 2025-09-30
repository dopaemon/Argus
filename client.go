package main

import (
	"context"
	"log"
	_ "time"
	"fmt"
	"bufio"
	"os"

	pb "example.com/chat/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// dùng insecure credentials cho local dev (thay bằng TLS trong production)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	ctx := context.Background()

	for {
		fmt.Print("Alice: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		res, err := c.SendMessage(ctx, &pb.MessageRequest{
			User: "Alice",
			Text: text,
		})
		if err != nil {
			log.Fatalf("SendMessage error: %v", err)
		}
		log.Printf("Server replied: %s", res.Reply)
	}
}
