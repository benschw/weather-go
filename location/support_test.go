package location

import (
	"fmt"
	"github.com/benschw/weather-go/openweather"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

func Test(t *testing.T) { TestingT(t) }

type TestWeatherClient struct {
}

func (c *TestWeatherClient) FindForLocation(city string, state string) (openweather.Conditions, error) {
	if city == "Austin" && state == "Texas" {
		return openweather.Conditions{
			Main: openweather.Main{
				Temperature: 75,
			},
			Weather: []openweather.Weather{
				openweather.Weather{
					Description: "sunny",
				},
			},
		}, nil
	} else {
		return openweather.Conditions{}, nil
	}
}
