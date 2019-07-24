package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKey := "7RBCSvtkuAj3NAFkgkGkJdvy6wssq4q1"

	c := goweather.NewClient(nil, apiKey)

	response, _, err := c.Location.GetCity("Bakirkoy")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
