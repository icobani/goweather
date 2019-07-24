// KAMİL KAPLAN

package test

import (
	"encoding/json"
	"github.com/icobani/goweather"
	"io/ioutil"
	"log"
	"testing"
)

func TestForeCast(t *testing.T) {
	apiKey := "7RBCSvtkuAj3NAFkgkGkJdvy6wssq4q1"
	isFile := false

	c := goweather.NewClient(nil, apiKey)
	var code string = "31822311"

	// Daha önce aranan code var ise onu code'a göre dosyaya kaydediyor. Eğer yok ise aranan
	// code.json dosyası oluşturuyor.
	files, err := ioutil.ReadDir("./weather")
	if err != nil {
		log.Fatal(err)
	}
	fileName := code + ".json"
	for _, f := range files {
		if f.Name() == fileName {
			data, err := ioutil.ReadFile("./weather/" + fileName)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("merhaba")
			log.Println(string(data))
			isFile = true
			break
		}
	}
	if !isFile {
		log.Println("asdasd")
		response, _, err := c.Forecast.GetForeCast(code)
		if err != nil {
			log.Fatal(err)
		}
		json_bytelar, _ := json.Marshal(response)
		err2 := ioutil.WriteFile("./weather/"+fileName, json_bytelar, 0777)
		if err2 != nil {
			// print it out
			log.Fatal(err2)
		}
		log.Println(response)
	}

}
