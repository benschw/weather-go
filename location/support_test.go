package location

import (
	"fmt"
	wapi "github.com/benschw/weather-go/openweather/api"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

func Test(t *testing.T) { TestingT(t) }

type TestWeatherClient struct {
}

func (c *TestWeatherClient) FindForLocation(city string, state string) (wapi.Conditions, error) {
	if city == "Austin" && state == "Texas" {
		return wapi.Conditions{
			Main: wapi.Main{
				Temperature: 75,
			},
			Weather: []wapi.Weather{
				wapi.Weather{
					Description: "sunny",
				},
			},
		}, nil
	} else {
		return wapi.Conditions{}, nil
	}
}
