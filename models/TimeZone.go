/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:45
   Notes   :
*/

package models

import "time"

type TimeZone struct {
	Code             string    `json:"code,omitempty"`
	Name             string    `json:"name,omitempty"`
	GmtOffset        float32   `json:"gmtOffset,omitempty"`
	IsDaylightSaving bool      `json:"ısDaylightSaving,omitempty"`
	NextOffsetChange time.Time `json:"nextOffsetChange,omitempty"`
}
