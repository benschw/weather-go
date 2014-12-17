package openweather

import (
	"fmt"
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

func (s *IntTestSuite) TestFind(c *C) {
	// given
	expectedId := 4671654
	expectedName := "Austin"
	client := &WeatherClient{}

	// when
	cond, err := client.FindForLocation("Austin", "Texas")

	// then
	c.Assert(err, Equals, nil)

	c.Assert(cond.Id, Equals, expectedId)
	c.Assert(cond.Name, Equals, expectedName)
}
