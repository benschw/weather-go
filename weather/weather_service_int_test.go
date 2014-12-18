package weather

import (
	"fmt"
	"github.com/benschw/opin-go/config"
	"github.com/benschw/opin-go/rando"
	"github.com/benschw/weather-go/weather/api"
	"github.com/benschw/weather-go/weather/client"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

func ARandomService(dbStr string) *WeatherService {
	host := fmt.Sprintf("localhost:%d", rando.Port())

	s := &WeatherService{
		Database: dbStr,
		Bind:     host,
	}
	go s.Run()

	return s
}

func Test(t *testing.T) { TestingT(t) }

type IntTestSuite struct {
	s    *WeatherService
	host string
}

var _ = Suite(&IntTestSuite{})

func (s *IntTestSuite) SetUpSuite(c *C) {
	var cfg struct {
		Database string
	}
	config.Bind("../test.yaml", &cfg)

	s.s = ARandomService(cfg.Database)
	s.host = "http://" + s.s.Bind

}
func (s *IntTestSuite) SetUpTest(c *C) {
	s.s.MigrateDb()
}
func (s *IntTestSuite) TearDownTest(c *C) {
	db, _ := s.s.getDb()
	db.DropTable(api.Location{})
}

func (s *IntTestSuite) TestAdd(c *C) {
	// given
	client := &client.LocationClient{Host: s.host}

	// when
	created, err := client.AddLocation("Austin", "Texas", 78751)

	// then
	c.Assert(err, Equals, nil)
	found, err := client.FindLocation(created.Id)

	c.Assert(created, DeepEquals, found)
}

func (s *IntTestSuite) TestFindAll(c *C) {
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

func (s *IntTestSuite) TestSave(c *C) {
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

func (s *IntTestSuite) TestDelete(c *C) {
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
