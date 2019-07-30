/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Forecast struct {
	LocationCode   string           `json:"LocationCode"`
	Headline       Headline         `json:"Headline"`
	DailyForecasts []DailyForecasts `json:"DailyForecasts"`
}
