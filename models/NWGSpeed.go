/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type NWGSpeed struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}
