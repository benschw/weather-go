package weather

import (
	"fmt"
	"github.com/benschw/opin-go/config"
	"github.com/benschw/opin-go/rando"
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

var _ = fmt.Print
var _ = log.Print

var cfg struct {
	Database string
}

var _ = config.Bind("../test.yaml", &cfg)

var server = ARandomService(cfg.Database)
var host = "http://" + server.Bind

func Test(t *testing.T) { TestingT(t) }

func ARandomService(dbStr string) *WeatherService {
	host := fmt.Sprintf("localhost:%d", rando.Port())

	s := &WeatherService{
		Database: dbStr,
		Bind:     host,
	}
	go s.Run()

	return s
}
