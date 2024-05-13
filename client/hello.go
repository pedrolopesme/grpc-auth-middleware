package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/pedrolopeme/grpc-auth-middleware/proto"
)

func doHello(client pb.HelloHandlerClient, name string) {
	fmt.Printf("doHello was invoked for %v\n", name)

	res, err := client.SayHello(context.Background(), &pb.Hello{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Response from server: %s", res.Response)
}
