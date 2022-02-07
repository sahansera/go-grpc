package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJson converts a protobuf message to a JSON string.
func ProtobufToJson(message proto.Message) ([]byte, error) {
	b := protojson.MarshalOptions{
		Indent:          "",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	return b.Marshal(message)
}

// JsonToProtobuf converts a JSON string to a protobuf message.
func JsonToProtobuf(data string, message proto.Message) error {
	b := protojson.UnmarshalOptions{
		AllowPartial: true,
	}
	return b.Unmarshal([]byte(data), message)
}
