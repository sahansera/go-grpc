package main

import (
	"context"
	"log"
	"net"

	errors2 "bookshop/errors"
	pb "bookshop/server/pb/inventory"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/cockroachdb/errors"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return nil, errors.Wrapf(errors.Wrapf(errors2.ErrBookshop, "second layer"), "third layer")
	//log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	//return &pb.GetBookListResponse{
	//	Books: getSampleBooks(),
	//}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
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
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}
