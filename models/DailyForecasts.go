/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type DailyForecasts struct {
	Date                     string `json:"date,omitempty"`
	EpochDate                int64  `json:"epochDate,omitempty"`
	Sun                      Sun
	Moon                     Moon
	Temperature              Temperature
	RealFeelTemperature      RealFeelTemperature
	RealFeelTemperatureShade RealFeelTemperatureShade
	HoursOfSun               float32 `json:"hoursOfSun,omitempty"`
	DegreeDaySummary         DegreeDaySummary
	AirAndPollen             []AirAndPollen
	Day                      Day
	Night                    Night
	Sources                  []string `json:"sources,omitempty"`
	MobieLink                string   `json:"mobieLink,omitempty"`
	Link                     string   `json:"link,omitempty"`
}
