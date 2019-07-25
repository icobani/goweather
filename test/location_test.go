package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKey := "fAySX5OL8h37IiIt6YKHSUyVQXtph9VM"

	c := goweather.NewClient(nil, apiKey)

	response, _, err := c.Location.GetCity("bakirkoy")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
