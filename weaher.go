package goweather

import (
	"encoding/json"
	"github.com/icobani/goweather/models"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = baseURL + "/locations/v1/cities/search"
	forecastURL = baseURL + "/forecasts/v1"
)

// A GoWather manages communication with the OneSignal API.
type GoWather struct {
	BaseURL      string                `json:"base_url"`
	LocationURL  string                `json:"location_url"`
	ForecastURL  string                `json:"forecast_url"`
	ApiKeys      string                `json:"api_key"`
	ApiKeyUsages []models.ApiKeyUsages `json:"api_key_usages"`
	Location     models.Location       `json:"location,omitempty"`
	ForeCast     models.Forecast       `json:"fore_cast,omitempty"`
}

type ErrorStruct struct {
	Error string
}

func (this GoWather) New(ApiKeys string, city string) (*ErrorStruct, *GoWather) {

	if ApiKeys == "" {
		return &ErrorStruct{"Api Key is cannot empty"}, nil
	}

	KeyUsages := SetApiKeys(ApiKeys)

	// Resty Default setup
	resty.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(3 * time.Second).
		SetTimeout(1 * time.Minute).
		SetContentLength(true)

	returnVal := &GoWather{
		ApiKeys:      ApiKeys,
		ApiKeyUsages: KeyUsages,
		BaseURL:      baseURL,
		LocationURL:  locationURL,
		ForecastURL:  forecastURL,
	}

	if city != "" {
		err := returnVal.SetLocation(city)
		if err != nil {
			return err, nil
		}
	}
	return nil, returnVal
}

func (this *GoWather) SetLocation(city string) *ErrorStruct {
	var savedlocations []models.Location

	fileIsExist, savedlocations := readLocations()
	if fileIsExist {
		// Sorun yok.
		for _, item := range savedlocations {
			if item.LocalizedName == city {
				this.Location = item
				return nil
			}
		}
	}

	// TODO : Eğer kayıtlarımızda yok ise apiden soracağız.
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey": this.GetApiKey(),
			"q":      city,
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.LocationURL)

	if err != nil {
		return &ErrorStruct{err.Error()}
	}

	var locations []models.Location
	err = json.Unmarshal(res.Body(), &locations)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	if len(locations) > 0 {
		this.Location = locations[0]
		if fileIsExist {
			// Demekki dosya var ama içinde aradığımmız zaman bizim location'ı bulamadık.
			// bu durumda dosyaya yeni bulduğumuz location'e da ekleyebiliriz.
			savedlocations = append(savedlocations, locations[0])
			errs := writeLocations(savedlocations)
			if errs != nil {
				return errs
			}
		} else {
			errs := writeLocations(locations)
			if errs != nil {
				return errs
			}
		}

	} else {
		return &ErrorStruct{"City name is missing"}
	}
	return nil
}

func fileIsExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readLocations() (bool, []models.Location) {
	var locations []models.Location
	var fileIsExist bool
	fileIsExist = fileIsExists("locations.json")
	if fileIsExist {
		file, _ := ioutil.ReadFile("locations.json")
		err := json.Unmarshal([]byte(file), &locations)

		if err != nil {
			return false, nil
		}
		return true, locations

	} else {
		return false, nil
	}
}

func readApiKeyUsages() (bool, []models.ApiKeyUsages) {
	var apiKeyUsages []models.ApiKeyUsages
	var fileIsExist bool
	fileIsExist = fileIsExists("apiKeyUsages.json")
	if fileIsExist {
		file, _ := ioutil.ReadFile("apiKeyUsages.json")
		err := json.Unmarshal([]byte(file), &apiKeyUsages)

		if err != nil {
			return false, nil
		}
		return true, apiKeyUsages

	} else {
		return false, nil
	}
}

func writeLocations(locations []models.Location) *ErrorStruct {
	var err = os.Remove("locations.json")
	fileBody, _ := json.MarshalIndent(locations, "", " ")
	err = ioutil.WriteFile("locations.json", fileBody, 0644)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	return nil
}

func writeApiKeyUsages(apiKeyUsages []models.ApiKeyUsages) *ErrorStruct {
	var err = os.Remove("apiKeyUsages.json")
	fileBody, _ := json.MarshalIndent(apiKeyUsages, "", " ")
	err = ioutil.WriteFile("apiKeyUsages.json", fileBody, 0644)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	return nil
}

func SetApiKeys(ApiKeys string) []models.ApiKeyUsages {
	var keys []string
	keys = strings.Split(ApiKeys, ",")
	fileIsExist, KeyUsages := readApiKeyUsages()
	if !fileIsExist {
		for _, item := range keys {
			KeyUsages = append(KeyUsages, models.ApiKeyUsages{
				ApiKey: item,
				Usage:  0,
			})
		}
		writeApiKeyUsages(KeyUsages)
	} else {
		var ItemFound bool
		var FileChanged bool
		for _, item := range keys {
			ItemFound = false
			for _, item2 := range KeyUsages {
				if item2.ApiKey == item {
					ItemFound = true
					break
				}
			}
			if !ItemFound {
				KeyUsages = append(KeyUsages, models.ApiKeyUsages{
					ApiKey: item,
					Usage:  0,
				})
				FileChanged = true
			}
		}
		if FileChanged {
			writeApiKeyUsages(KeyUsages)
		}
	}
	return KeyUsages
}

func (this *GoWather) GetApiKey() string {
	sort.Slice(this.ApiKeyUsages, func(i, j int) bool {
		return this.ApiKeyUsages[i].Usage < this.ApiKeyUsages[j].Usage
	})
	this.ApiKeyUsages[0].Usage += 1
	writeApiKeyUsages(this.ApiKeyUsages)
	return this.ApiKeyUsages[0].ApiKey

}
