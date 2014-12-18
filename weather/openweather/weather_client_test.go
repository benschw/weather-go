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
	// when
	cond, err := FindForLocation("Austin", "Texas")

	// then
	c.Assert(err, Equals, nil)

	c.Assert(cond.Main.Temperature > 0, Equals, true)
	c.Assert(cond.Weather[0].Description, Not(Equals), "")
}
