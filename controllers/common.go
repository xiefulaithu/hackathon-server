package controllers

// BasicResp defines the basic response format
type BasicResp struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
