package wetter

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const requesturl string = "https://api.brightsky.dev/weather?"

type FallBackSourceId struct {
	PressureMSL   int32  `json:"pressure_msl" example:"4"`
	Visibility    int32  `json:"visibility" example:"0"`
	WindSpeed     int32  `json:"wind_speed" example:"3"`
	WindGustSpeed int32  `json:"wind_gust_speed" example:"5"`
	CloudCover    int32  `json:"cloud_cover" example:"2"`
	WindDirection int32  `json:"wind_direction" example:"1"`
	Icon          string `json:"icon" example:"cloud"`
}

type WeatherPoint struct {
	Timestamp         string           `json:"timestamp" example:"2020-10-25T00:00:00+00:00"`
	SourceId          float32          `json:"source_id" example:"6946"`
	Precipitation     float32          `json:"precipitation,omitempty" example:"1007.6"`
	PressureMSL       float32          `json:"pressure_msl,omitempty" example:"14.1"`
	Sunshine          float32          `json:"sunshine,omitempty" example:"200"`
	Temperature       float32          `json:"temperature,omitempty" example:"19.4"`
	WindDirection     float32          `json:"wind_direction,omitempty" example:"100"`
	WindSpeed         float32          `json:"wind_speed,omitempty" example:"13.5"`
	CloudCover        float32          `json:"cloud_cover,omitempty" example:"90"`
	DewPoint          float32          `json:"dew_point,omitempty" example:"10.6"`
	RelativeHumidity  float32          `json:"relative_humidity,omitempty" example:"31.2"`
	Visibility        float32          `json:"visibility,omitempty" example:"41.2"`
	WindGustDirection float32          `json:"wind_gust_direction,omitempty" example:"40.2"`
	WindGustSpeed     float32          `json:"wind_gust_speed,omitempty" example:"12.2"`
	Condition         string           `json:"condition,omitempty" example:"dry"`
	FallBackSourceIds FallBackSourceId `json:"fallback_source_ids"`
}

type WeatherData struct {
	Weather []WeatherPoint `json:"weather"`
}

func Weather(lan, lon string, date time.Time) (WeatherData, int, error) {
	request_url := getWeatherURL(lan, lon, date.Format(time.RFC3339))

	fmt.Println(request_url)

	var weather WeatherData

	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Println(err)
		return weather, http.StatusBadRequest, err
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, &weather)
	if err != nil {
		log.Fatalln(err)
	}

	return weather, resp.StatusCode, nil

}

func getWeatherURL(lat, lon, date string) string {
	return fmt.Sprintf("%slat=%s&lon=%s&date=%s", requesturl, lat, lon, date)
}
