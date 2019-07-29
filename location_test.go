package goweather

import (
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKeys := "fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu"
	var err, goWeather = GoWather{}.New(apiKeys, "İstanbul", "beşiktaş")
	if err != nil {
		log.Fatal(err)
	} else {
		date := goWeather.ForeCast.Headline.EffectiveDate
		log.Print("Date: ", date)
		log.Printf("GoWeather : %v\n", goWeather)
	}
}
