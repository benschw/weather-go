package main

import (
	"github.com/benschw/weather-go/location"
	"log"
	"os"
)

func main() {
	bindAddress := os.Getenv("bind")
	databaseDsn := os.Getenv("database")

	log.Printf("bindAddress: %s dsn: %s", bindAddress, databaseDsn)

	s, err := location.NewLocationService(bindAddress, databaseDsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.MigrateDb(); err != nil {
		log.Fatal(err)
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
