/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:44
   Notes   :
*/

package models

type Country struct {
	ID            string `json:"id,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
}
