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

	dummyAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"

	c := pb.NewHelloHandlerClient(conn)
	doHello(c, "Pedro", dummyAccessToken)

}
