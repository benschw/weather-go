package client

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"github.com/benschw/weather-go/location/api"
	"log"
	"net/http"
)

var _ = log.Print

type LocationClient struct {
	Host string
}

func (c *LocationClient) AddLocation(city string, state string, zipcode int) (api.Location, error) {
	var location api.Location

	newLocation := api.Location{
		City:    city,
		State:   state,
		Zipcode: zipcode,
	}

	url := fmt.Sprintf("%s/location", c.Host)
	r, err := rest.MakeRequest("POST", url, newLocation)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusCreated)
	return location, err
}

func (c *LocationClient) FindAllLocations() ([]api.Location, error) {
	var locations []api.Location

	url := fmt.Sprintf("%s/location", c.Host)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}
	err = rest.ProcessResponseEntity(r, &locations, http.StatusOK)
	return locations, err
}

func (c *LocationClient) FindLocation(id int) (api.Location, error) {
	var location api.Location

	url := fmt.Sprintf("%s/location/%d", c.Host, id)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusOK)
	return location, err
}

func (c *LocationClient) SaveLocation(toSave api.Location) (api.Location, error) {
	var location api.Location

	url := fmt.Sprintf("%s/location/%d", c.Host, toSave.Id)
	r, err := rest.MakeRequest("PUT", url, toSave)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusOK)
	return location, err
}

func (c *LocationClient) DeleteLocation(id int) error {
	url := fmt.Sprintf("%s/location/%d", c.Host, id)
	r, err := rest.MakeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	err = rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
	return err
}
