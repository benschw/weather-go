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

type TestSuite struct {
	s    *WeatherService
	host string
}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpSuite(c *C) {
	s.s = server
	s.host = host
}
func (s *TestSuite) SetUpTest(c *C) {
	s.s.MigrateDb()
}
func (s *TestSuite) TearDownTest(c *C) {
	db, _ := s.s.getDb()
	db.DropTable(api.Location{})
}

func (s *TestSuite) TestAdd(c *C) {
	// given
	client := &client.LocationClient{Host: s.host}

	// when
	created, err := client.AddLocation("Austin", "Texas", 78751)

	// then
	c.Assert(err, Equals, nil)
	found, err := client.FindLocation(created.Id)

	c.Assert(created, DeepEquals, found)
}

func (s *TestSuite) TestFindAll(c *C) {
	// given
	client := &client.LocationClient{Host: s.host}

	loc1, err := client.AddLocation("Austin", "Texas", 78751)
	loc2, err := client.AddLocation("Williamsburg", "Virginia", 23188)
	// when

	foundLocations, err := client.FindAllLocations()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(foundLocations, DeepEquals, []api.Location{loc1, loc2})
}

func (s *TestSuite) TestSave(c *C) {
	// given
	client := &client.LocationClient{Host: s.host}
	location, err := client.AddLocation("Austin", "Texas", 78751)

	// when
	location.State = "foo"
	saved, err := client.SaveLocation(location)

	// then
	c.Assert(err, Equals, nil)

	c.Assert(location.State, DeepEquals, saved.State)
}

func (s *TestSuite) TestDelete(c *C) {
	// given
	client := &client.LocationClient{Host: s.host}
	location, err := client.AddLocation("Austin", "Texas", 78751)

	// when
	err = client.DeleteLocation(location.Id)

	// then
	c.Assert(err, Equals, nil)

	foundLocations, _ := client.FindAllLocations()

	c.Assert(len(foundLocations), Equals, 0)
}
