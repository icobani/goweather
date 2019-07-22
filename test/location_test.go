package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKey := "EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs"

	c := goweather.NewClient(nil, apiKey)

	response, _, err := c.Location.GetCity("istanbul")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
