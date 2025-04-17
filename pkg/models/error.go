package models

type ErrorFound struct {
	Message   string `json:"message"`
	Error     string `json:"error"`
	Timestamp string `json:"timestamp"`
}
