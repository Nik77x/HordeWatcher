package data

type ModelStatus struct {
	Name   string  `json:"name"`
	Queued float64 `json:"queued"`
	Count  int     `json:"count"`
}
