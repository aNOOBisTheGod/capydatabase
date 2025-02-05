package utils

var (
	NUMBER_OF_IMAGES int
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ImageStruct struct {
	URL    string `json:"url"`
	Index  int    `json:"index"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
