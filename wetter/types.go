package wetter

import "time"

type GeoJSONPoint struct {
	Type        string     `json:"type"`        // point type is "Point"
	Coordinates [2]float64 `json:"coordinates"` // coordinates are lon lat
}

type WeatherRequest struct {
	Date         time.Time    `json:"date,omitempty"`      // z time string
	LastDate     time.Time    `json:"last_date,omitempty"` // z time string
	Geopoint     GeoJSONPoint `json:"point,omitempty"`     // geo point lon lat
	DwdStationId string       `json:"dwd_id,omitempty"`
	WmoStationId string       `json:"wmo_id,omitempty"`
	MaxDist      int          `json:"max_dist,omitempty"` // Meter
}

type WeatherResponse struct {
	Timestamp         time.Time `json:"timestamp"`
	SourceId          int32     `json:"source_id"`
	Precipitation     float32   `json:"precipitation"`       // kg/m^3
	Presure           int32     `json:"pressure_msl"`        // Pascal
	Sunshine          int64     `json:"sunshine"`            // seconds
	Temperature       float32   `json:"temperature"`         // Kerlvin
	WindDirection     int32     `json:"wind_direction"`      // ° Degree
	WindSpeed         float32   `json:"wind_speed"`          // m/s
	CloudCover        int8      `json:"cloud_cover"`         // %
	DewPoint          float32   `json:"dew_point"`           // Kelvin
	RelativeHumidity  int8      `json:"relative_humidity"`   // %
	Visibility        int64     `json:"visibility"`          // Meter
	WindGustDirection int16     `json:"wind_gust_direction"` // ° Degree
	WindGustSpeed     float32   `json:"wind_gust_speed"`     // m/s
	Condition         string    `json:"condition"`
	Icon              string    `json:"icon"`
	Valid             bool
}
type WeatherResponses struct {
	Weather []WeatherResponse
}

type SourceResponse struct {
	Id              int64     `json:"id"`
	DwdStationId    string    `json:"dwd_station_id"`
	ObservationType string    `json:"observation_type"`
	Lat             float64   `json:"lat"`
	Lon             float64   `json:"lon"`
	Height          float32   `json:"height"`
	StationName     string    `json:"station_name"`
	WmoStationId    string    `json:"wmo_station_id"`
	FirstRecord     time.Time `json:"first_record"`
	LastRecord      time.Time `json:"last_record"`
	Distance        float64   `json:"distance"`
}

type FallBackSource struct {
	WindSpeed         int32 `json:"wind_speed"`
	WindDirection     int32 `json:"wind_direction"`
	WindGustSpeed     int32 `json:"wind_gust_speed"`
	WindGustDirection int32 `json:"wind_gust_direction"`
	PressureMSL       int32 `json:"pressure_msl"`
	Visibility        int32 `json:"visibility"`
	CloudCover        int32 `json:"cloud_cover"`
}

type DWDWeather struct {
	Timestamp         time.Time      `json:"timestamp"`
	SourceId          int32          `json:"source_id"`
	Precipitation     float32        `json:"precipitation"`       // kg/m^3
	Presure           int32          `json:"pressure_msl"`        // Pascal
	Sunshine          int64          `json:"sunshine"`            // seconds
	Temperature       float32        `json:"temperature"`         // Kerlvin
	WindDirection     int32          `json:"wind_direction"`      // ° Degree
	WindSpeed         float32        `json:"wind_speed"`          // m/s
	CloudCover        int8           `json:"cloud_cover"`         // %
	DewPoint          float32        `json:"dew_point"`           // Kelvin
	RelativeHumidity  int8           `json:"relative_humidity"`   // %
	Visibility        int64          `json:"visibility"`          // Meter
	WindGustDirection int16          `json:"wind_gust_direction"` // ° Degree
	WindGustSpeed     float32        `json:"wind_gust_speed"`     // m/s
	Condition         string         `json:"condition"`
	Icon              string         `json:"icon"`
	FallbackSourceID  FallBackSource `json:"fallback_source_ids"`
}

type WeatherDWDResponse struct {
	Weather []DWDWeather     `json:"weather"`
	Sources []SourceResponse `json:"sources"`
}

type DWDSourcesResponse struct {
	Source []SourceResponse `json:"sources"`
}
