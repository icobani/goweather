package goweather

import (
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKeys := "Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,JbuSGqPN2WrQvChrPxx8vXjd1ekYshQh,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs0,sJXIyoAmMS68m2ACUWuAxGnzEp0OGcv0,fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu,yLxyo5AWCoO8vHOLphOty0QsFxj5sUe2,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,tdfew2mhPXGw3rs18ACrS1DHk5mMGg49"
	var err, goWeather = GoWather{}.New(apiKeys, "Malatya", "Kale")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v\n", goWeather)
	}
}
