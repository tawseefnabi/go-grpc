package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tawseefnabi/go-grpc.git/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen on port 9091: %v", err)
	}
	fmt.Println("server running at port: 9001")
	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc on port 9091: %v", err)
	}

}
