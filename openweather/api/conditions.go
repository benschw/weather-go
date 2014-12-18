package api

type Conditions struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

type Weather struct {
	Description string `json:"description"`
}

type Main struct {
	Temperature float32 `json:"temp"`
}
