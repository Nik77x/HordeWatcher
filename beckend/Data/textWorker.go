package data

type TextWorker struct {
	Worker
	// text data
	MaxLength        int `json:"max_length"`
	MaxContextLength int `json:"max_context_length"`
}
