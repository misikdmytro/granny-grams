package model

type ImageGenerationRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      uint8  `json:"n"`
	Size   string `json:"size"`
}

type ImageGenerationData struct {
	URL string `json:"url"`
}

type ImageGenerationResponse struct {
	Data []ImageGenerationData `json:"data"`
}

type OpenAPIErrorData struct {
	Code    *string `json:"code"`
	Message string  `json:"message"`
	Param   *string `json:"param"`
	Type    string  `json:"type"`
}

type OpenAPIErrorResponse struct {
	Error OpenAPIErrorData `json:"error"`
}
