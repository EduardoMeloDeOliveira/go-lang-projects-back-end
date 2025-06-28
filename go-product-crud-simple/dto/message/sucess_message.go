package message

type SuccessMessage struct {
	Message  string `json : "message"`
	StatusCode int64 `json : staus_code`
}