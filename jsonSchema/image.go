package jsonSchema

type ImageModel string

const OpenAiDalle2 ImageModel = "OpenAiDalle2"
const OpenAiDalle3 ImageModel = "OpenAiDalle3"

type ImageSize string

const (
	//this code is nicked from go-openai
	CreateImageSize256x256   ImageSize = "256x256"
	CreateImageSize512x512   ImageSize = "512x512"
	CreateImageSize1024x1024 ImageSize = "1024x1024"

	// dall-e-3 supported only.
	CreateImageSize1792x1024 ImageSize = "1792x1024"
	CreateImageSize1024x1792 ImageSize = "1024x1792"
)

// Image if you want the Url of the image use the DataType String otherwise use the DataType Byte
type Image struct {
	Model ImageModel `json:"model,omitempty"`
	Size  ImageSize  `json:"size,omitempty"`
}
