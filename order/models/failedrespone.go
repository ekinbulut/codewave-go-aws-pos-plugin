package models

type FailedResponse struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}
