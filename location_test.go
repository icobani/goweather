package goweather

import (
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKeys := "sJXIyoAmMS68m2ACUWuAxGnzEp0OGcv0,JbuSGqPN2WrQvChrPxx8vXjd1ekYshQh,Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu,yLxyo5AWCoO8vHOLphOty0QsFxj5sUe2,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs0"
	var err, goWeather = GoWather{}.New(apiKeys, "İstanbul", "Fatih")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v\n", goWeather)
	}
}
