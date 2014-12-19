package location

import (
	wapi "github.com/benschw/weather-go/openweather/api"
)

type WeatherClient interface {
	FindForLocation(city string, state string) (wapi.Conditions, error)
}
