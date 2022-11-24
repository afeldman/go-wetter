package wetter

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const sourcerequesturl string = "https://api.brightsky.dev/sources?units=si&tz=UTC"

func Source(source_id int32) {
	// build url
	//lat, lon float64, date, last_date, dwd_station, wmo_station string, source, max_dist int
	request_url := fmt.Sprintf("%s")

	fmt.Println(request_url)
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		log.Println(err)
		return
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

	log.Println(string(b))

	/*err = json.Unmarshal(b, &weather)
	if err != nil {
		log.Fatalln(err)
	}*/

	//return wr, resp.StatusCode, nil

}
