package api

type Location struct {
	Id          int     `json:"id"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Temperature float32 `json:"temperature"`
	Weather     string  `json:"weather"`
}
