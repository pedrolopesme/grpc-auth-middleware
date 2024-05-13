package main

import (
	pb "github.com/pedrolopeme/grpc-auth-middleware/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const server_addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(server_addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloHandlerClient(conn)
	doHello(c, "Pedro")

}
