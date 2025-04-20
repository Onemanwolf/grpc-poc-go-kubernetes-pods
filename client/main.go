package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "go.client.grpc/pb"
)

func main() {
	conn, err := grpc.Dial("server-service:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Send 10 requests
	for i := 1; i <= 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
		if err != nil {
			log.Printf("Request %d failed: %v", i, err)
			continue
		}
		log.Printf("Request %d - Greeting: %s", i, r.GetMessage())

		// Delay between requests
		time.Sleep(1 * time.Second)
	}

	log.Println("Completed 10 requests, shutting down")
}