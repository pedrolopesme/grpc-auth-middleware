package main

import (
	"log"
	"net"

	pb "github.com/pedrolopeme/grpc-auth-middleware/proto"
	"google.golang.org/grpc"
)

const addr = "localhost:50051"

type Server struct {
	pb.HelloHandlerServer
}

func main() {
	// Create a listener on the specified address.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	// Log the server's listening address.
	log.Printf("server listening at %v", listener.Addr())

	// Create a new gRPC server.
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			AuthInterceptor,
		),
	)

	// Register the HelloHandlerServer implementation with the server.
	pb.RegisterHelloHandlerServer(server, &Server{})

	// Start serving requests on the listener.
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
