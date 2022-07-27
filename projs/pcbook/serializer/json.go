package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJSON(msg proto.Message) (string, error) {
	return (&jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
	}).MarshalToString(msg)
}
