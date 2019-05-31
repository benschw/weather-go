package main

import (
	"flag"
	"fmt"
	"github.com/benschw/weather-go/location"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	cmd := flag.Arg(0)

	bindAddress := os.Getenv("bind")
	databaseDsn := os.Getenv("database")

	// Configure Server
	s, err := location.NewLocationService(bindAddress, databaseDsn)
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
	case "help":
		flag.Usage()
	case "migrate-db":
		if err := s.MigrateDb(); err != nil {
			log.Fatal(err)
		}
	default:
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	}

}
