package main

import (
	"context"
	"log"

	pb "bookshop/client/pb/inventory"
	errors2 "bookshop/errors"

	"github.com/cockroachdb/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	_, err = client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		if errors.Is(err, errors2.ErrBookshop) {
			log.Println("RECOGNIZED")
		} else {
			log.Println("NOT RECOGNIZED")
		}
		log.Printf("CAUSE: %v", errors.Cause(err))
	}
	//log.Printf("book list: %v", bookList)
}
