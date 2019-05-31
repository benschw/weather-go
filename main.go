package main

import (
	"flag"
	"fmt"
	"github.com/benschw/opin-go/config"
	"github.com/benschw/weather-go/location"
	"log"
	"os"
)

func main() {
	// Get Arguments
	var cfgPath string

	flag.StringVar(&cfgPath, "config", "./config.yaml", "Path to Config File")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Load Config
	var cfg struct {
		Bind     string
		Database string
	}
	if err := config.Bind(cfgPath, &cfg); err != nil {
		log.Fatal(err)
	}

	// pull desired command/operation from args
	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatal("Command argument required")
	}
	cmd := flag.Arg(0)

	// Configure Server
	s, err := location.NewLocationService(cfg.Bind, cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Run Main App
	switch cmd {
	case "serve":

		// Start Server
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	case "migrate-db":

		// Start Server
		if err := s.MigrateDb(); err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		log.Fatalf("Unknown Command: %s", cmd)
	}

}
