/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:47
   Notes   :
*/

package models

type Details struct {
	Key                      string `json:"key,omitempty"`
	StationCode              string `json:"stationCode,omitempty"`
	StationGmtOffset         int64  `json:"stationGmtOffset,omitempty"`
	BandMap                  string `json:"bandMap,omitempty"`
	Climo                    string `json:"climo,omitempty"`
	LocalRadar               string `json:"localRadar,omitempty"`
	MediaRegion              string `json:"mediaRegion,omitempty"`
	Metar                    string `json:"metar,omitempty"`
	NXMetro                  string `json:"nxMetro,omitempty"`
	NXState                  string `json:"nxState,omitempty"`
	Population               int64  `json:"population,omitempty"`
	PrimaryWarningCountyCode string `json:"primaryWarningCountyCode,omitempty"`
	PrimaryWarningZoneCode   string `json:"primaryWarningZoneCode,omitempty"`
	Satellite                string `json:"satellite,omitempty"`
	Synoptic                 string `json:"synoptic,omitempty"`
	MarineStation            string `json:"marineStation,omitempty"`
	MarineStationGMTOffset   int64  `json:"marineStationGmtOffset,omitempty"`
	VideoCode                string `json:"videoCode,omitempty"`
	PartnerID                int32  `json:"partnerId,omitempty"`
	Sources                  []Sources
	CanonicalPostalCode      string `json:"canonicalPostalCode,omitempty"`
	CanonicalLocationKey     string `json:"canonicalLocationKey,omitempty"`
	LocationStem             string `json:"locationStem,omitempty"`
}
