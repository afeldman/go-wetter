package wetter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const sourcerequesturl string = "https://api.brightsky.dev/sources?"

func Source(c *gin.Context, source_id int32) {
	// build url
	//lat, lon float64, date, last_date, dwd_station, wmo_station string, source, max_dist int
	request_url := fmt.Sprintf("%ssource_id=%d", sourcerequesturl, source_id)

	client := &http.Client{}
	resp, err := client.Get(request_url)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	var dwdresp DWDSourcesResponse

	err = json.Unmarshal(b, &dwdresp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, dwdresp.Source)

}
