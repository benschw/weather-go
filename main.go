package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/benschw/weather-go/weather"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Bind     string
	Database string
}

func LoadConfig(path string) (Config, error) {
	config := Config{}

	if _, err := os.Stat(path); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

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
	cfg, err := LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// pull desired command/operation from args
	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatal("Command argument required")
	}
	cmd := flag.Arg(0)

	// Configure Server
	s := &weather.Server{
		Database: cfg.Database,
		Bind:     cfg.Bind,
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
		if err := s.Migrate(); err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		log.Fatalf("Unknown Command: %s", cmd)
	}

}
