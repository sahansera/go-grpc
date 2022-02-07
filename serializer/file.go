package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToJSONFile writes a protobuf message to a JSON file.
func WriteProtobufToJSONFile(filePath string, message proto.Message) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("failed to marshal proto message: %v", err)
	}

	err = ioutil.WriteFile(filePath, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// WriteProtobufToBinaryFile writes a protobuf message to a binary file.
func WriteProtobufToBinaryFile(filePath string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal proto message: %v", err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// ReadProtobufFromBinaryFile reads a protobuf message from a binary file.
func ReadProtobufFromBinaryFile(filePath string, message proto.Message) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal proto message: %v", err)
	}

	return nil
}
