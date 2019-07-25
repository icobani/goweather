/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:45
   Notes   :
*/

package models

type AdministrativeArea struct {
	ID            string `json:"id,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
	Level         int32  `json:"level,omitempty"`
	LocalizedType string `json:"localizedType,omitempty"`
	EnglishType   string `json:"englishType,omitempty"`
	CountryID     string `json:"countryId,omitempty"`
}
