package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/pedrolopeme/grpc-auth-middleware/proto"
	"google.golang.org/grpc/metadata"
)

func doHello(client pb.HelloHandlerClient, name string, accessToken string) {
	fmt.Printf("doHello was invoked for %v\n", name)

	// Add the access token to the request metadata.
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", accessToken))

	// building the message
	message := &pb.Hello{Name: name}

	// Sending message
	res, err := client.SayHello(ctx, message)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Response from server: %s", res.Response)
}
