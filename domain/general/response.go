package general

type ResponseData struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type ResponseMessageData struct {
	Message string `json:"message,omitempty"`
}
