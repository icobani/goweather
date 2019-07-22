package goweather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = "http://dataservice.accuweather.com/locations/v1"
	forecastURL = "http://localhost:1478"
)

type Client struct {
	BaseURL     *url.URL
	LocationURL *url.URL
	ForecastURL *url.URL
	ApiKey      string
	Client      *http.Client

	Location *LocationService
	Forecast *ForecastService
}

type ErrorResponse struct {
	Messages []string `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	msg := "Weather returned those error messages:\n - "
	return msg + strings.Join(e.Messages, "\n - ")
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	locationURL, err := url.Parse(locationURL)
	if err != nil {
		log.Fatal(err)
	}

	forecastURL, err := url.Parse(forecastURL)
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		BaseURL:     baseURL,
		LocationURL: locationURL,
		ForecastURL: forecastURL,
		Client:      httpClient,
		ApiKey:      apiKey,
	}

	c.Location = &LocationService{client: c}
	c.Forecast = &ForecastService{client: c}

	return c
}

func (c *Client) NewRequest(method string, t string, path string) (*http.Request, error) {
	wUrl := ""
	switch t {
	case "location":
		wUrl = c.LocationURL.String()
		break
	case "forecast":
		wUrl = c.ForecastURL.String()
		break
	default:
		wUrl = c.BaseURL.String()
		break
	}

	u, err := url.Parse(wUrl + path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return &ErrorResponse{
			Messages: []string{"Internal Server Error"},
		}
	default:
		var errResp ErrorResponse
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&errResp)
		if err != nil {
			errResp.Messages = []string{"Couldn't decode response body JSON"}
		}
		return &errResp
	}
}
