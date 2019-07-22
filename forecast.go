package goweather

type ForecastService struct {
	client *Client
}

type Forecast struct {
	Headline Headline
	DailyForecasts []DailyForecasts
}

type Headline struct {
	EffectiveDate      string `json:"effective_date"`
	EffectiveEpochDate int64  `json:"effective_epoch_date"`
	Severity           int32  `json:"severity"`
	Text               string `json:"text"`
	Category           string `json:"category"`
	EndDate            string `json:"end_date"`
	EndEpochDate       int64  `json:"end_epoch_date"`
	MobileLink         string `json:"mobile_link"`
	Link               string `json:"link"`
}

type DailyForecasts struct {

}