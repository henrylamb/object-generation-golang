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
		SpeechToText: &jsonSchema.SpeechToText{
			Model:             (*jsonSchema.SpeechToTextModel)(&protoDef.SpeechToText.Model),
			AudioToTranscribe: protoDef.SpeechToText.AudioToTranscribe,
			Language:          protoDef.SpeechToText.Language,
			ToString:          protoDef.SpeechToText.ToString,
			ToCaptions:        protoDef.SpeechToText.ToCaptions,
			Format:            jsonSchema.AudioFormat(protoDef.SpeechToText.Format),
		},
		TextToSpeech: &jsonSchema.TextToSpeech{
			Model:         (*jsonSchema.TextToSpeechModel)(&protoDef.TextToSpeech.Model),
			Voice:         jsonSchema.Voice(protoDef.TextToSpeech.Voice),
			StringToAudio: protoDef.TextToSpeech.StringToAudio,
			Format:        jsonSchema.AudioFormat(protoDef.TextToSpeech.Format),
		},
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
		Image: &pb.Image{
			Model: string(*modelDef.Image.Model),
			Size:  string(*modelDef.Image.Size),
		},
		SpeechToText: &pb.SpeechToText{
			Model:             string(*modelDef.SpeechToText.Model),
			AudioToTranscribe: modelDef.SpeechToText.AudioToTranscribe,
			Language:          modelDef.SpeechToText.Language,
			ToString:          modelDef.SpeechToText.ToString,
			ToCaptions:        modelDef.SpeechToText.ToCaptions,
			Format:            string(modelDef.SpeechToText.Format),
		},
		TextToSpeech: &pb.TextToSpeech{
			Model:         string(*modelDef.TextToSpeech.Model),
			StringToAudio: modelDef.TextToSpeech.StringToAudio,
			Format:        string(modelDef.TextToSpeech.Format),
			Voice:         string(modelDef.TextToSpeech.Voice),
		},
	}

	for key, modelProperty := range modelDef.Properties {
		protoDef.Properties[key] = ConvertModelToProto(&modelProperty)
	}

	return protoDef
}
