package weather

import (
	"fmt"
	"github.com/benschw/weather-go/weather/api"
	"github.com/benschw/weather-go/weather/client"
	. "gopkg.in/check.v1"
	"log"
)

var _ = fmt.Print
var _ = log.Print

type IntTestSuite struct {
	s    *WeatherService
	host string
}

var _ = Suite(&IntTestSuite{})

func (s *IntTestSuite) SetUpSuite(c *C) {
	s.s = server
	s.host = host
}
func (s *IntTestSuite) SetUpTest(c *C) {
	s.s.MigrateDb()
}
func (s *IntTestSuite) TearDownTest(c *C) {
	db, _ := s.s.getDb()
	db.DropTable(api.Location{})
}

func (s *IntTestSuite) TestAdd(c *C) {

	// when
	created, err := client.AddLocation(s.host, "Austin", "Texas", 78751)

	// then
	c.Assert(err, Equals, nil)
	found, err := client.FindLocation(s.host, created.Id)

	c.Assert(created, DeepEquals, found)
}

func (s *IntTestSuite) TestFindAll(c *C) {
	// given
	loc1, err := client.AddLocation(s.host, "Austin", "Texas", 78751)
	loc2, err := client.AddLocation(s.host, "Williamsburg", "Virginia", 23188)
	// when

	foundLocations, err := client.FindAllLocations(s.host)

	// then
	c.Assert(err, Equals, nil)

	c.Assert(foundLocations, DeepEquals, []api.Location{loc1, loc2})
}

func (s *IntTestSuite) TestSave(c *C) {
	// given
	location, err := client.AddLocation(s.host, "Austin", "Texas", 78751)

	// when
	location.State = "foo"
	saved, err := client.SaveLocation(s.host, location)

	// then
	c.Assert(err, Equals, nil)

	c.Assert(location.State, DeepEquals, saved.State)
}

func (s *IntTestSuite) TestDelete(c *C) {
	// given
	location, err := client.AddLocation(s.host, "Austin", "Texas", 78751)

	// when
	err = client.DeleteLocation(s.host, location.Id)

	// then
	c.Assert(err, Equals, nil)

	foundLocations, _ := client.FindAllLocations(s.host)

	c.Assert(len(foundLocations), Equals, 0)
}
