/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Direction struct {
	Degress   float64 `json:"degress,omitempty"`
	Localized string  `json:"localized,omitempty"`
	English   string  `json:"english,omitempty"`
}
