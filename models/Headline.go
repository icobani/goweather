/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Headline struct {
	EffectiveDate      string `json:"effectiveDate,omitempty"`
	EffectiveEpochDate int64  `json:"effectiveEpochDate,omitempty"`
	Severity           int32  `json:"severity,omitempty"`
	Text               string `json:"text,omitempty"`
	Category           string `json:"category,omitempty"`
	EndDate            string `json:"endDate,omitempty"`
	EndEpochDate       int64  `json:"endEpochDate,omitempty"`
	MobileLink         string `json:"mobileLink,omitempty"`
	Link               string `json:"link,omitempty"`
}