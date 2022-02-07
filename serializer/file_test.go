package serializer

import (
	"main/mocks"
	"main/pb"
	"testing"
)

func TestFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/booklist_msg.bin"
	bookList1 := mocks.NewBookList()

	// Write the book list to a binary file.
	if err := WriteProtobufToBinaryFile(binaryFile, bookList1); err != nil {
		t.Fatalf("failed to write binary file: %v", err)
	}

	tmpBookList := pb.BookList{}
	err := ReadProtobufFromBinaryFile(binaryFile, &tmpBookList)

	if err != nil {
		t.Fatalf("failed to read binary file")
	}

	if tmpBookList.Books[0].Title != bookList1.Books[0].Title {
		t.Fatalf("failed to read binary file")
	}

}
