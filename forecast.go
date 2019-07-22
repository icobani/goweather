package goweather

import (
	"net/http"
	"net/url"
)

type ForecastService struct {
	client *Client
}

type Forecast struct {
	Headline Headline
	// DailyForecasts []DailyForecasts
}

type Headline struct {
	EffectiveDate      string `json:"effectiveDate"`
	EffectiveEpochDate int64  `json:"effectiveEpochDate"`
	Severity           int32  `json:"severity"`
	Text               string `json:"text"`
	Category           string `json:"category"`
	EndDate            string `json:"end_date"`
	EndEpochDate       int64  `json:"end_epoch_date"`
	MobileLink         string `json:"mobile_link"`
	Link               string `json:"link"`
}

/*type DailyForecasts struct {
	Date string `json:"date"`
}*/

func (f *ForecastService) GetForeCast() (*Headline, *http.Response, error) {
	// path := fmt.Sprintf("/cities/search?apikey=%s&q=%s", l.client.ApiKey, city)
	//pull request denemesi
	path := "/apiTest"

	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	req, err := f.client.NewRequest("GET", "forecast", u.String())
	if err != nil {
		return nil, nil, err
	}

	lResp := new(Headline)
	resp, err := f.client.Do(req, lResp)
	if err != nil {
		return nil, resp, err
	}
	return lResp, resp, err
}
