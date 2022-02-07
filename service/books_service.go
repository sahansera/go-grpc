package service

import (
	"context"
	"main/pb"
)

type BookList struct {
}

func NewBooksService() *BookList {
	return &BookList{}
}

func GetBooks(ctx context.Context, in *pb.GetBookListRequest) []*pb.Book {
	return []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
			Language:  "en",
		},
		{
			Title:     "2001: A Space Odyssey",
			Author:    "Arthur C. Clarke",
			PageCount: 100,
			Language:  "en",
		},
	}
}
