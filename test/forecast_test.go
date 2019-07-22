package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestForeCast(t *testing.T) {
	c := goweather.NewClient(nil, "")

	response, _, err := c.Forecast.GetForeCast()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

}
