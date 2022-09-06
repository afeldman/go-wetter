package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	wetter "github.com/afeldman/go-wetter/wetter"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "wetter/docs"

	"github.com/gin-gonic/gin"
)

var (
	release bool
	port    int
)

func init() {
	flag.BoolVar(&release, "release", false, "set to release mode")
	flag.IntVar(&port, "port", 2510, "set the port")
}

func local_context(ctx *gin.Context, lon, lat string, date time.Time) {
	var weatherresponse WeatherResponse
	data, code, err := wetter.Weather(lat, lon, date)
	if err != nil {
		weatherresponse.Status = http.StatusInternalServerError
		weatherresponse.Error = err.Error()
		weatherresponse.WeatherData = nil

		ctx.JSON(http.StatusInternalServerError, weatherresponse)
	}

	weatherresponse.Status = code
	weatherresponse.Error = ""
	weatherresponse.WeatherData = &data
	ctx.JSON(code, weatherresponse)
}

type WeatherResponse struct {
	Status int `json:"status" example:"404"`
	*wetter.WeatherData
	Error string `json:"error,omitempty" example:"cannot load data"`
}

// @Summary      get weather information by date
// @Description  get a weather description of dwd
// @Param 		 date	query	string	true "2020-10-25"
// @Param		 lat	query	string 	true "51.873960"
// @Param		 lon	query	string  true "8.156710"
// @Success      200  {object}  WeatherResponse
// @Failure      500  {object}  WeatherResponse
// @Failure      404  {object}  WeatherResponse
// @Router		 /{date}/{lat}/{lon} [get]
func weather_by_date(ctx *gin.Context) {
	lat := ctx.Param("lat")
	lon := ctx.Param("lon")
	sdate := ctx.Param("date")

	var weatherresponse WeatherResponse

	date, error := time.Parse("2006-01-02", sdate)
	if error != nil {
		weatherresponse.Status = http.StatusInternalServerError
		weatherresponse.WeatherData = nil
		weatherresponse.Error = error.Error()
		ctx.JSON(http.StatusInternalServerError, weatherresponse)
	}

	local_context(ctx, lon, lat, date)

}

// @Summary      Show an account
// @Description  get string by ID
// @Param		 lat	query	string 	true "51.873960"
// @Param		 lon	query	string  true "8.156710"
// @Success      200  {object}  WeatherResponse
// @Failure      500  {object}  WeatherResponse
// @Failure      404  {object}  WeatherResponse
// @Router		 /now/{lat}/{lon} [get]
func weather_now(ctx *gin.Context) {
	lat := ctx.Param("lat")
	lon := ctx.Param("lon")

	local_context(ctx, lon, lat, time.Now())
}

// @Title          Weather information download
// @Version        0.1.0
// @Description    request for weather informations
// @TermsOfService http://.../terms

// @Contact.name  Weather API
// @Contact.url   http://weather.....com
// @Contact.email anton.feldmann@gmail.com

// @License.name MIT
// @License.url  https://MIT-license.org

// @Host     http://weather....com
// @BasePath /v1
func main() {

	flag.Parse()

	router := gin.New()

	if release {
		log.Println("start release mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("start development mode")
		gin.SetMode(gin.DebugMode)
	}

	v1 := router.Group("/v1")
	{
		v1.GET("/:date/:lat/:lon", weather_by_date)

		v1.GET("/now/:lat/:lon", weather_now)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
