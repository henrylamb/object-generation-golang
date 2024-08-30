package converison

import (
	pb "github.com/henrylamb/object-generation-golang/grpc"
	"github.com/henrylamb/object-generation-golang/jsonSchema"
)

// ConvertProtoToModel converts a protobuf Definition to your Go model Definition
func ConvertProtoToModel(protoDef *pb.Definition) *jsonSchema.Definition {
	if protoDef == nil {
		return nil
	}

	modelDef := &jsonSchema.Definition{
		Type:               jsonSchema.DataType(protoDef.Type),
		Instruction:        protoDef.Instruction,
		Properties:         make(map[string]jsonSchema.Definition),
		Required:           protoDef.Required,
		Items:              ConvertProtoToModel(protoDef.Items),
		Model:              jsonSchema.ModelType(protoDef.Model),
		ProcessingOrder:    protoDef.ProcessingOrder,
		SystemPrompt:       &protoDef.SystemPrompt,
		ImprovementProcess: protoDef.ImprovementProcess,
		SelectFields:       protoDef.SelectFields,
		Voters:             protoDef.Voters,
		HashMap:            ConvertProtoToHashMap(protoDef.HashMap),
		NarrowFocus:        ConvertProtoToFocus(protoDef.NarrowFocus),
		Req:                ConvertProtoToRequestFormat(protoDef.Req),
		Choices:            ConvertProtoToChoices(protoDef.Choices),
	}

	for key, protoProperty := range protoDef.Properties {
		modelDef.Properties[key] = *ConvertProtoToModel(protoProperty)
	}

	return modelDef
}

// ConvertModelToProto converts your Go model Definition to a protobuf Definition
func ConvertModelToProto(modelDef *jsonSchema.Definition) *pb.Definition {
	if modelDef == nil {
		return nil
	}

	systemPrompt := ""
	if modelDef.SystemPrompt != nil {
		systemPrompt = *modelDef.SystemPrompt
	}

	protoDef := &pb.Definition{
		Type:               string(modelDef.Type),
		Instruction:        modelDef.Instruction,
		Properties:         make(map[string]*pb.Definition),
		Required:           modelDef.Required,
		Items:              ConvertModelToProto(modelDef.Items),
		Model:              string(modelDef.Model),
		ProcessingOrder:    modelDef.ProcessingOrder,
		SystemPrompt:       systemPrompt,
		ImprovementProcess: modelDef.ImprovementProcess,
		SelectFields:       modelDef.SelectFields,
		Voters:             modelDef.Voters,
		HashMap:            ConvertModelToProtoHashMap(modelDef.HashMap),
		NarrowFocus:        ConvertModelToProtoFocus(modelDef.NarrowFocus),
		Req:                ConvertModelToProtoRequestFormat(modelDef.Req),
		Choices:            ConvertModelToProtoChoices(modelDef.Choices),
	}

	for key, modelProperty := range modelDef.Properties {
		protoDef.Properties[key] = ConvertModelToProto(&modelProperty)
	}

	return protoDef
}