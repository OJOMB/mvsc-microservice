package utils

// ApplicationError models the response object returned on error
type ApplicationError struct {
	Msg    string `json:"message"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
