package main

import (
	"log"
	"main/pb"

	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBookListServiceServer
}

func main() {
	// Start gRPC server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookListServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
