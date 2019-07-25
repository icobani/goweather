/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:44
   Notes   :
*/

package models

type Location struct {
	Version                int32  `json:"Version,omitempty"`
	Key                    string `json:"key,omitempty"`
	Type                   string `json:"type,omitempty"`
	Rank                   int32  `json:"rank,omitempty"`
	LocalizedName          string `json:"LocalizedName,omitempty"`
	EnglishName            string `json:"EnglishName,omitempty"`
	PrimaryPostalCode      string `json:"PrimaryPostalCode,omitempty"`
	Region                 Regions
	Country                Country
	AdministrativeArea     AdministrativeArea
	TimeZone               TimeZone
	GeoPosition            GeoPosition
	IsAlias                bool `json:"ıs_alias,omitempty"`
	SupplementalAdminAreas []SupplementalAdminAreas
	DataSets               []string `json:"data_sets,omitempty"`
	Details                Details
}
