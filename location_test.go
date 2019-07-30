package goweather

import (
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKeys := "fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu"
	// distict değerinin baş harfi büyük olmak zorunda eğer büyük harfle başlamazsa apiden çeker var olan değeri
	var err, goWeather = GoWather{}.New(apiKeys, "İstanbul", "Fatih")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v\n", goWeather)
	}
}
