package domain

type Response struct {
	Code   int32       `json:"code,omitempty"`
	Object interface{} `json:"object,omitempty"`
	Err    error       `json:"err,omitempty"`
	Text   string      `json:"text,omitempty"`
}
