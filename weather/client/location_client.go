package client

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"github.com/benschw/weather-go/weather/api"
	"log"
	"net/http"
)

var _ = log.Print

var AddLocation = func(host string, city string, state string, zipcode int) (api.Location, error) {
	var location api.Location

	newLocation := api.Location{
		City:    city,
		State:   state,
		Zipcode: zipcode,
	}

	url := fmt.Sprintf("%s/location", host)
	r, err := rest.MakeRequest("POST", url, newLocation)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusCreated)
	return location, err
}

var FindAllLocations = func(host string) ([]api.Location, error) {
	var locations []api.Location

	url := fmt.Sprintf("%s/location", host)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return locations, err
	}
	err = rest.ProcessResponseEntity(r, &locations, http.StatusOK)
	return locations, err
}

var FindLocation = func(host string, id int) (api.Location, error) {
	var location api.Location

	url := fmt.Sprintf("%s/location/%d", host, id)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusOK)
	return location, err
}

var SaveLocation = func(host string, toSave api.Location) (api.Location, error) {
	var location api.Location

	url := fmt.Sprintf("%s/location/%d", host, toSave.Id)
	r, err := rest.MakeRequest("PUT", url, toSave)
	if err != nil {
		return location, err
	}
	err = rest.ProcessResponseEntity(r, &location, http.StatusOK)
	return location, err
}

var DeleteLocation = func(host string, id int) error {
	url := fmt.Sprintf("%s/location/%d", host, id)
	r, err := rest.MakeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	err = rest.ProcessResponseEntity(r, nil, http.StatusNoContent)
	return err
}
