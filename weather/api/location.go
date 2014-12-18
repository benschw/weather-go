package api

import (
	"github.com/benschw/weather-go/weather/openweather"
)

type Location struct {
	Id         int                    `json:"id"`
	City       string                 `json:"city"`
	State      string                 `json:"state"`
	Zipcode    int                    `json:"zipcode"`
	Conditions openweather.Conditions `json:"conditions"`
}
