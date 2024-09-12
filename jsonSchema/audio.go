package jsonSchema

// AudioModel constant types for the different auido models
type TextToSpeechModel string

const OpenAiWhisper TextToSpeechModel = "OpenAiWhisper"
const GroqWhisper TextToSpeechModel = "GroqWhisper"

// TextToSpeech the DataType to use with this type is Byte
type TextToSpeech struct {
	Model         *TextToSpeechModel `json:"model,omitempty"`
	StringToAudio string             `json:"stringToAudio,omitempty"`
}

type SpeechToTextModel string

const OpenAiTTS SpeechToTextModel = "tts"

// SpeechToText the DataType to use with this type is String
type SpeechToText struct {
	Model             *SpeechToTextModel `json:"model,omitempty"`
	AudioToTranscribe []byte             `json:"audioToTranscribe,omitempty"`

	//This defines how the audio will be returned. Either in the format of captions or just a string
	ToCaptions bool `json:"captions,omitempty"`
	ToString   bool `json:"toString,omitempty"`
}
