package models

type Response struct {
	RemoteResponse RemoteResponse `json:"remoteResponse"`
}

type RemoteResponse struct {
	RemoteOrderID string `json:"remoteOrderId"`
}
