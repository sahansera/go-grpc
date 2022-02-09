package main

import (
	"bookshop/server/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleBooks() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
	}
	return sampleBooks
}
