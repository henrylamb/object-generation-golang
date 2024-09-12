package jsonSchema

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

// AudioModel constant types for the different auido models
type AudioModel string

const OpenAiWhisper AudioModel = "OpenAiWhisper"
const GroqWhisper AudioModel = "GroqWhisper"

// Audio
// the prompt information for the image will exist in the instructions like the other types
// the audio type object will need to be present along with the selection of the audio type being used for this content to be interaction within the generation on rails
// the audio will only support the create of the audio and not the processing of an audio file into a transcript (for now)
// this information will be returned as a string as a base64 encoded string
type Audio struct {
	Model *AudioModel `json:"model,omitempty"`
	Data  []byte      `json:"-"` // used for transmitting an audio file
}

// MarshalJSON custom JSON marshaller to encode Data as base64
func (a *Audio) MarshalJSON() ([]byte, error) {
	type Alias Audio
	return json.Marshal(&struct {
		*Alias
		Base64Content string `json:"base64Content,omitempty"`
	}{
		Alias: (*Alias)(a),
		//This means that the information will need an unmarshal function for the audio data
		Base64Content: base64.StdEncoding.EncodeToString(a.Data),
	})
}

// The processing of the Audio file could be a format where if the value is filled in the audio struct then it leads to that being processed into a transcript and that transcript would be
//set as the string value for the return type and

// UnmarshalJSON custom JSON unmarshaller to decode base64 to Data
func (a *Audio) UnmarshalJSON(data []byte) error {
	type Alias Audio
	aux := &struct {
		*Alias
		Base64Content string `json:"base64Content,omitempty"`
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.Base64Content != "" {
		decodedData, err := base64.StdEncoding.DecodeString(aux.Base64Content)
		if err != nil {
			return errors.New("failed to decode base64 content")
		}
		a.Data = decodedData
	}

	return nil
}
