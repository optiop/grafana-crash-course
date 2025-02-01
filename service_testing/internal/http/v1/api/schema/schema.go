package schema

type Message struct {
	Message string `json:"message"`
	Status  string `json:"error"` // panic, fail, error, success
}
