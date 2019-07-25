/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:46
   Notes   :
*/

package models

type SupplementalAdminAreas struct {
	Level         int32  `json:"level,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
}
