package goweather

import (
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKey := "fAySX5OL8h37IiIt6YKHSUyVQXtph9VM"
	var err, goWeather = GoWather{}.New(apiKey, "İstanbul")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v", goWeather)
	}
}
