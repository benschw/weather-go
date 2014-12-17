package api

type Location struct {
	Id      int    `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}
