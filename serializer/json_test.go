package serializer

import (
	"main/mocks"
	"testing"
)

func TestJson(t *testing.T) {
	t.Parallel()

	jsonFile := "../tmp/booklist_msg.json"
	bookList1 := mocks.NewBookList()

	// Write the book list to a JSON file.
	if err := WriteProtobufToJSONFile(jsonFile, bookList1); err != nil {
		t.Fatalf("failed to write JSON file: %v", err)
	}
}
