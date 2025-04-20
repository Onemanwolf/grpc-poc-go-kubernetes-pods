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

	// Counter for request numbering
	requestNum := 1

	// Run indefinitely, sending a request every 10 seconds
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
		if err != nil {
			log.Printf("Request %d failed: %v", requestNum, err)
		} else {
			log.Printf("Request %d - Greeting: %s", requestNum, r.GetMessage())
		}
		cancel()

		// Increment request counter
		requestNum++

		// Wait 10 seconds before the next request
		time.Sleep(10 * time.Second)
	}
}