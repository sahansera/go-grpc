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

	tmpBookList := &pb.BookList{}
	ReadProtobufFromBinaryFile(binaryFile, tmpBookList)

	if tmpBookList == nil {
		t.Fatalf("failed to read binary file")
	}

}
