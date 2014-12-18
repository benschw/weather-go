package openweather

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"log"
	"net/http"
)

var _ = log.Print

const UriString string = "http://api.openweathermap.org/data/2.5/weather?units=imperial&q=" //Austin,Texas

type WeatherClient struct {
}

func (c *WeatherClient) FindForLocation(city string, state string) (Conditions, error) {
	var cond Conditions

	url := fmt.Sprintf("%s%s,%s", UriString, city, state)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return cond, err
	}
	err = rest.ProcessResponseEntity(r, &cond, http.StatusOK)
	return cond, err
}
