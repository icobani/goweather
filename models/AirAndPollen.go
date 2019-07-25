/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type AirAndPollen struct {
	Name          string `json:"name,omitempty"`
	Value         int32  `json:"value,omitempty"`
	Category      string `json:"category,omitempty"`
	CategoryValue int32  `json:"categoryValue,omitempty"`
	Type          string `json:"type,omitempty"`
}
