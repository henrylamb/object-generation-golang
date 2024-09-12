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

type Image struct {
	Model *ImageModel `json:"model,omitempty"`
	Size  *ImageSize  `json:"format,omitempty"`
	//the below determine how the data is returned back to the user. ie how the string is returned back to them and how it should be processed
	Url    bool `json:"url,omitempty"` //if this is true then the url returned will be format of /generated/"image url"
	Base64 bool `json:"base64,omitempty"`
	//return format --> this is the format that the image will be returned in

}
