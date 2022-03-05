package models

type FailureResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}
