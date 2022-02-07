package mocks

import "main/pb"

// NewBookList creates a new BookList with sample data.
func NewBookList() *pb.BookList {
	return &pb.BookList{
		Books: []*pb.Book{
			{
				Title:     "Book 1",
				Author:    "Author 1",
				PageCount: 20,
				Language:  "en",
			},
			{
				Title:     "Book 2",
				Author:    "Author 2",
				PageCount: 200,
				Language:  "en",
			},
			{
				Title:     "Book 3",
				Author:    "Author 4",
				PageCount: 2000,
				Language:  "en",
			},
		},
	}
}
