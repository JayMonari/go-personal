package serializer

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToJSONFile(msg proto.Message, filename string) error {
	data, err := ProtobufToJSON(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}
	if err = os.WriteFile(filename, []byte(data), 0644); err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}

// WriteProtobufToBinaryFile writes protocol buffer message to named file.
func WriteProtobufToBinaryFile(msg proto.Message, filename string) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	if err = os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile reads a protocol buffer message from binary file.
func ReadProtobufFromBinaryFile(filename string, msg proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}
	if err = proto.Unmarshal(data, msg); err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}
