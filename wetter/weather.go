package wetter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const weatherrequesturl string = "https://api.brightsky.dev/weather?units=si&tz=UTC"

func NewWeatherResponces() *WeatherResponses {
	return &WeatherResponses{
		Weather: []WeatherResponse{},
	}
}

func Weather(wr WeatherRequest, ch chan<- WeatherResponses) {
	// if the date is none, set current time for start
	if wr.Date.IsZero() {
		wr.Date = time.Now()
	}
	// if there is no current time or the next date is
	// before the current date add 24 hours to the current date
	// to get the next date in line
	if wr.LastDate.IsZero() || wr.LastDate.Before(wr.Date) {
		// add one day to date if date not set
		wr.LastDate = wr.Date.Add(time.Hour)
	}

	wr.DwdStationId = strings.TrimSpace(wr.DwdStationId)
	wr.WmoStationId = strings.TrimSpace(wr.WmoStationId)

	if wr.MaxDist <= 0 {
		wr.MaxDist = 50_000 // m
	}

	// check for lat, lon, dwd or wmo
	if wr.DwdStationId == "" &&
		wr.WmoStationId == "" &&
		wr.Geopoint.Coordinates[0] == 0 &&
		wr.Geopoint.Coordinates[1] == 0 {
		log.Println("need point")
		return
	}

	// build url
	request_url := getWeatherURL(&wr)

	client := &http.Client{}
	resp, err := client.Get(request_url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var weathers WeatherDWDResponse

	err = json.Unmarshal(b, &weathers)
	if err != nil {
		log.Println(err.Error())
		return
	}

	wartherresp := WeatherResponses{}
	wartherresp.Weather = make([]WeatherResponse, len(weathers.Weather))

	for index, weather := range weathers.Weather {
		weatherrep := WeatherResponse{
			Timestamp:         weather.Timestamp,
			SourceId:          weather.SourceId,
			Precipitation:     weather.Precipitation,
			Presure:           weather.Presure,
			Sunshine:          weather.Sunshine,
			Temperature:       weather.Temperature,
			WindDirection:     weather.WindDirection,
			WindSpeed:         weather.WindSpeed,
			CloudCover:        weather.CloudCover,
			DewPoint:          weather.DewPoint,
			RelativeHumidity:  weather.RelativeHumidity,
			Visibility:        weather.Visibility,
			WindGustDirection: weather.WindGustDirection,
			WindGustSpeed:     weather.WindGustSpeed,
			Condition:         weather.Condition,
			Icon:              weather.Icon,
			Valid:             true,
		}
		wartherresp.Weather[index] = weatherrep
	}

	ch <- wartherresp
}

func getWeatherURL(wr *WeatherRequest) string {
	base := fmt.Sprintf("%s&date=%s&last_date=%s&max_dist=%d",
		weatherrequesturl,
		wr.Date.Format("2006-01-02"),
		wr.LastDate.Format("2006-01-02"),
		wr.MaxDist)

	if wr.DwdStationId != "" {
		base += fmt.Sprintf("&dwd_station_id=%s", wr.DwdStationId)
	} else if wr.WmoStationId != "" {
		base += fmt.Sprintf("&wmo_station_id=%s", wr.WmoStationId)
	} else {
		base += fmt.Sprintf("&lon=%f&lat=%f", wr.Geopoint.Coordinates[0], wr.Geopoint.Coordinates[1])
	}

	return base
}
