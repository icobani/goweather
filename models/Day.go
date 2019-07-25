/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Day struct {
	Icon                     int32  `json:"ıcon,omitempty"`
	IconPhrase               string `json:"ıconPhrase,omitempty"`
	HasPrecipitation         bool   `json:"hasPrecipitation,omitempty"`
	LocalSource              DLocalSource
	ShortPhrase              string `json:"shortPhrase,omitempty"`
	LongPhrase               string `json:"longPhrase,omitempty"`
	PrecipitationProbability int32  `json:"precipitationProbability,omitempty"`
	ThunderstormProbability  int32  `json:"thunderstormProbability,omitempty"`
	RainProbability          int32  `json:"rainProbability,omitempty"`
	SnowProbability          int32  `json:"snowProbability,omitempty"`
	IceProbability           int32  `json:"ıceProbability,omitempty"`
	Wind                     Wind
	WindGust                 WindGust
	TotalLiquid              TotalLiquid
	Rain                     Rain
	Snow                     Snow
	Ice                      Ice
	HoursOfPrecipitation     float32 `json:"hoursOfPrecipitation,omitempty"`
	HoursOfRain              float32 `json:"hoursOfRain,omitempty"`
	HoursOfSnow              float32 `json:"hoursOfSnow,omitempty"`
	HoursOfIce               float32 `json:"hoursOfIce,omitempty"`
	CloudCover               int32   `json:"cloudCover,omitempty"`
}