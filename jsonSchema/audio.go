package jsonSchema

// AudioModel constant types for the different auido models
type TextToSpeechModel string

const OpenAiTTS TextToSpeechModel = "tts"

// TextToSpeech the DataType to use with this type is Byte
type TextToSpeech struct {
	Model         *TextToSpeechModel `json:"model,omitempty"`
	StringToAudio string             `json:"stringToAudio,omitempty"`
	//This defines how the audio will be returned. Either in the format of captions or just a string
	//TODO need to alter all the GRPC and converision logic to handle this
	Voice  Voice       `json:"voice,omitempty"`
	Format AudioFormat `json:"format,omitempty"`
}

type SpeechToTextModel string

const OpenAiWhisper SpeechToTextModel = "OpenAiWhisper"
const GroqWhisper SpeechToTextModel = "GroqWhisper"

type AudioFormat string

const Text AudioFormat = "text"
const SRT AudioFormat = "srt"
const VTT AudioFormat = "vtt"
const JSON AudioFormat = "json"
const VerboseJSON AudioFormat = "verbose-json"

type Voice string

const (
	Alloy   Voice = "alloy"
	Echo    Voice = "echo"
	Fable   Voice = "fable"
	Onyx    Voice = "onyx"
	Nova    Voice = "nova"
	Shimmer Voice = "shimmer"
)

// SpeechToText the DataType to use with this type is String
type SpeechToText struct {
	Model             *SpeechToTextModel `json:"model,omitempty"`
	AudioToTranscribe []byte             `json:"audioToTranscribe,omitempty"`
	Language          string             `json:"language,omitempty"` //must be in the format of ISO-639-1  will default to en (english)
	Format            AudioFormat        `json:"format,omitempty"`
	ToString          bool               `json:"toString,omitempty"`
	ToCaptions        bool               `json:"toCaptions,omitempty"`
}
