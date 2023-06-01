package model

type ParamUploadImage struct {
	Base64StringImage string `json:"base_64_string_image"`
	ApiKey            string `json:"api_key"`
}

type ImageInfo struct {
	Image struct {
		Url string `json:"url"`
	} `json:"image"`
}
