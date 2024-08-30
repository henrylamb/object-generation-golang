package converison

import (
	pb "github.com/henrylamb/object-generation-golang/grpc"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
)

// ConvertProtoToChoices converts a protobuf Choices to the Go model Choices
func ConvertProtoToChoices(protoChoices *pb.Choices) *jsonSchema.Choices {
	if protoChoices == nil {
		return nil
	}

	return &jsonSchema.Choices{
		Number:  int(protoChoices.Number),
		Options: protoChoices.Options,
	}
}

// ConvertModelToProtoChoices converts a Go model Choices to a protobuf Choices
func ConvertModelToProtoChoices(modelChoices *jsonSchema.Choices) *pb.Choices {
	if modelChoices == nil {
		return nil
	}

	return &pb.Choices{
		Number:  int32(modelChoices.Number),
		Options: modelChoices.Options,
	}
}
