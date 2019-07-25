package goweather

import (
	"encoding/json"
	"github.com/icobani/goweather/models"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
	"time"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = baseURL + "/locations/v1"
	forecastURL = baseURL + "/forecasts/v1"
)

// A GoWather manages communication with the OneSignal API.
type GoWather struct {
	BaseURL     string          `json:"base_url"`
	LocationURL string          `json:"location_url"`
	ForecastURL string          `json:"forecast_url"`
	ApiKey      string          `json:"api_key"`
	Location    models.Location `json:"location,omitempty"`
	ForeCast    models.Forecast `json:"fore_cast,omitempty"`
}

type ErrorStruct struct {
	Error string
}

func (this GoWather) New(ApiKey string, city string) (*ErrorStruct, *GoWather) {

	if ApiKey == "" {
		return &ErrorStruct{"Api Key is cannot empty"}, nil
	}

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
		ApiKey:      ApiKey,
		BaseURL:     baseURL,
		LocationURL: locationURL,
		ForecastURL: forecastURL,
	}

	if city != "" {
		returnVal.SetLocation(city)
	}
	return nil, returnVal
}

func (this *GoWather) SetLocation(city string) *ErrorStruct {

	// TODO : Eğer kayıtlarımızda yok ise apiden soracağız.
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey": this.ApiKey,
			"q":      city,
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.LocationURL)

	if err != nil {
		return &ErrorStruct{err.Error()}
	}

	var result []models.Location
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	if len(result) > 0 {
		this.Location = result[0]

		if !FileIsExists("locations.json") {
			file, _ := json.MarshalIndent(result, "json", " ")
			_ = ioutil.WriteFile("locations.json", file, 0644)
		}

	} else {
		return &ErrorStruct{"City name is missing"}
	}
	return nil
}

func FileIsExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
