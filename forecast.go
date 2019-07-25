package goweather

import (
	"net/http"
	"net/url"
)

// /forecasts/v1/daily/1day/318251?apikey=EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs
func (f *ForecastService) GetForeCast(code string) (*Forecast, *http.Response, error) {
	path := "/daily/1day/" + code + "?apikey=" + f.client.ApiKey
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	req, err := f.client.NewRequest("GET", "forecast", u.String())
	if err != nil {
		return nil, nil, err
	}

	fResp := new(Forecast)
	resp, err := f.client.Do(req, fResp)
	if err != nil {
		return nil, resp, err
	}

	/*fah := fResp.DailyForecasts[0].Temperature.Minimum.Value
	cel := fah2cel(fah)
	fResp.DailyForecasts[0].Temperature.Minimum.Value = cel
	fResp.DailyForecasts[0].Temperature.Minimum.Unit  = "C"*/

	return fResp, resp, err
}
