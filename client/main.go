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

	dummyAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJleGFtcGxlLmNvbSIsInN1YiI6IjEyMzQ1Njc4OTAiLCJhdWQiOiJleGFtcGxlLmNvbSIsImV4cCI6MzI0NzIxMTUyMCwiaWF0IjoxNjIyNTQ4ODAwLCJzY29wZSI6WyJkZW1vIl19.RdwHG0ANpFLoQQhnAxLgyn40TRR50R1Tz0FevloIDZY"

	c := pb.NewHelloHandlerClient(conn)
	doHello(c, "Pedro", dummyAccessToken)

}
