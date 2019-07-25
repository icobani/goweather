/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:44
   Notes   :
*/

package models

import (
	"gopkg.in/resty.v1"
	"log"
	"net/http"
	"time"
)

type Locations struct {
	Version                int32  `json:"version,omitempty"`
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

func (this *Locations) GetLocation(LocationName string) {

	uri := "http://dataservice.accuweather.com/locations/v1/cities/search?apikey=" + l.client.ApiKey + "&q=" + LocationName
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"limit": "10000",
		}).
		SetAuthToken(h.AccessToken).
		SetHeader("Content-Type", "application/json").
		Get(uri)

	if err != nil {
		return err
	}

	if res.StatusCode() != http.StatusOK {
		log.Println(res.Status())
		log.Println(res)
		return errors.New(12, "Bir sorun ile karşılaşıldı.")
	}

	result := InHouseListResponse{}
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return err
	}
}
