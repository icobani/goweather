package goweather

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

type LocationService struct {
	client *Client
}

type Locations struct {
	Version                int32  `json:"version"`
	Key                    string `json:"key"`
	Type                   string `json:"type"`
	Rank                   int32  `json:"rank"`
	LocalizedName          string `json:"LocalizedName"`
	EnglishName            string `json:"EnglishName"`
	PrimaryPostalCode      string `json:"PrimaryPostalCode"`
	Region                 Regions
	Country                Country
	AdministrativeArea     AdministrativeArea
	TimeZone               TimeZone
	GeoPosition            GeoPosition
	IsAlias                bool `json:"ıs_alias"`
	SupplementalAdminAreas []SupplementalAdminAreas
	DataSets               []string `json:"data_sets"`
	Details                Details
}

type Regions struct {
	ID            string `json:"id"`
	LocalizedName string `json:"LocalizedName"`
	EnglishName   string `json:"englishName"`
}

type Country struct {
	ID            string `json:"id"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
}

type AdministrativeArea struct {
	ID            string `json:"id"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
	Level         int32  `json:"level"`
	LocalizedType string `json:"localizedType"`
	EnglishType   string `json:"englishType"`
	CountryID     string `json:"countryId"`
}

type TimeZone struct {
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	GmtOffset        float32   `json:"gmtOffset"`
	IsDaylightSaving bool      `json:"ısDaylightSaving"`
	NextOffsetChange time.Time `json:"nextOffsetChange"`
}

type GeoPosition struct {
	Latitude  float64
	Longitude float64
	Elevation Elevation
}

type Elevation struct {
	Metric   Metric
	Imperial Imperial
}

type Metric struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	UnitType int32   `json:"unitType"`
}

type Imperial struct {
	Value    float64 `json:"value"`
	Unit     string  `json:"unit"`
	UnitType int32   `json:"unitType"`
}

type SupplementalAdminAreas struct {
	Level         int32  `json:"level"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
}

type Details struct {
	Key                      string `json:"key"`
	StationCode              string `json:"stationCode"`
	StationGmtOffset         int64  `json:"stationGmtOffset"`
	BandMap                  string `json:"bandMap"`
	Climo                    string `json:"climo"`
	LocalRadar               string `json:"localRadar"`
	MediaRegion              string `json:"mediaRegion"`
	Metar                    string `json:"metar"`
	NXMetro                  string `json:"nxMetro"`
	NXState                  string `json:"nxState"`
	Population               int64  `json:"population"`
	PrimaryWarningCountyCode string `json:"primaryWarningCountyCode"`
	PrimaryWarningZoneCode   string `json:"primaryWarningZoneCode"`
	Satellite                string `json:"satellite"`
	Synoptic                 string `json:"synoptic"`
	MarineStation            string `json:"marineStation"`
	MarineStationGMTOffset   int64  `json:"marineStationGmtOffset"`
	VideoCode                string `json:"videoCode"`
	PartnerID                int32  `json:"partnerId"`
	Sources                  []Sources
	CanonicalPostalCode      string `json:"canonicalPostalCode"`
	CanonicalLocationKey     string `json:"canonicalLocationKey"`
	LocationStem             string `json:"locationStem"`
}

type DMA struct {
	ID          string `json:"id"`
	EnglishName string `json:"englishName"`
}

type Sources struct {
	DataType string `json:"dataType"`
	Source   string `json:"source"`
	SourceId int32  `json:"sourceId"`
}

func (l *LocationService) GetCity(city string) (*[]Locations, *http.Response, error) {
	// path := fmt.Sprintf("/cities/search?apikey=%s&q=%s", l.client.ApiKey, city)

	path := "/cities/search?apikey=" + l.client.ApiKey + "&q=" + city
	log.Println(path)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	req, err := l.client.NewRequest("GET", "location", u.String())
	if err != nil {
		return nil, nil, err
	}

	lResp := new([]Locations)
	resp, err := l.client.Do(req, lResp)
	if err != nil {
		return nil, resp, err
	}
	return lResp, resp, err
}
