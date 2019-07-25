/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:46
   Notes   :
*/

package models

type Metric struct {
	Value    float64 `json:"value ,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}
