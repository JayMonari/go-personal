package serializer_test

import (
	"grpbook/pb"
	"grpbook/sample"
	"grpbook/serializer"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	laptop1 := sample.NewLaptop()
	require.NoError(t, serializer.WriteProtobufToBinaryFile(
		laptop1, "/tmp/laptop.bin"))

	laptop2 := &pb.Laptop{}
	require.NoError(t, serializer.ReadProtobufFromBinaryFile(
		"/tmp/laptop.bin", laptop2))
	require.True(t, proto.Equal(laptop1, laptop2))

	require.NoError(t, serializer.WriteProtobufToJSONFile(
		laptop1, "/tmp/laptop.json"))
}
