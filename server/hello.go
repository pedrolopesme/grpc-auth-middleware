package main

import (
	"context"
	"fmt"

	pb "github.com/pedrolopeme/grpc-auth-middleware/proto"
)

func (s *Server) SayHello(ctx context.Context, in *pb.Hello) (*pb.HelloResponse, error) {
	fmt.Printf("Hello was invoked with %v\n", in)

	response := &pb.HelloResponse{
		Response: fmt.Sprintf("Hello %s", in.Name),
	}

	return response, nil
}
