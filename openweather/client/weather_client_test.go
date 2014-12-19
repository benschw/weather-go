package client

import (
	"fmt"
	"github.com/benschw/weather-go/openweather/api"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

func Test(t *testing.T) { TestingT(t) }

type IntTestSuite struct {
}

var _ = Suite(&IntTestSuite{})

// Find should return weather for a city/state
func (s *IntTestSuite) TestFind(c *C) {
	// given
	client := WeatherClient{}

	// when
	cond, err := client.FindForLocation("Austin", "Texas")

	// then
	c.Assert(err, Equals, nil)

	c.Assert(cond.Main.Temperature > 0, Equals, true)
	c.Assert(cond.Weather[0].Description, Not(Equals), "")
}

// Client should return empty "Conditions" when a state isn't found
func (s *IntTestSuite) TestFindNotFound(c *C) {
	// given
	client := WeatherClient{}

	// when
	cond, err := client.FindForLocation("Foo", "Bar")

	// then
	c.Assert(err, Equals, nil)

	c.Assert(cond, DeepEquals, api.Conditions{})
}
