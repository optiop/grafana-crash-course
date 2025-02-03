package schema

type Message struct {
	Message string `json:"message"`
	Status  string `json:"status"` // panic, error, fail, success
}
