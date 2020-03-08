package models

type Response struct {
	Type       string `json:"type"`
	Value      string `json:"value"`
	StatusCode int    `json:"status_code"`
}

type Rule struct {
	Path     string `json:"path"`
	Response `json:"response"`
}
