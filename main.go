package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "wetter/docs"
	"wetter/wetter"

	"github.com/gin-gonic/gin"
)

var (
	release bool
	port    int
	server  string
)

func init() {
	flag.BoolVar(&release, "release", false, "set to release mode")
	flag.IntVar(&port, "port", 2510, "set the port")
	flag.StringVar(&server, "server", "0.0.0.0", "set server address")
}

// @Summary      Show an account
// @Description  get string by ID
// @Success      200  {object}  WeatherResponse
// @Failure      500  {object}  WeatherResponse
// @Failure      404  {object}  WeatherResponse
// @Router		 / [post]
func weather(c *gin.Context) {
	// read the body to json string
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	// request a list of weather data
	var requests []wetter.WeatherRequest

	// json string to request object
	if err := json.Unmarshal([]byte(jsonData), &requests); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// build channel for parallel request
	responseChannel := make(chan wetter.WeatherResponses)

	for _, req := range requests {
		go wetter.Weather(req, responseChannel)
	}

	response := make([][]wetter.WeatherResponse, 0)
	for range requests {
		k := <-responseChannel
		response = append(response, k.Weather)
	}

	c.JSON(http.StatusOK, response)

}

func source(c *gin.Context) {
	c.Param("source_id")
	c.JSON(http.StatusOK, "cool geht :)")
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
		v1.POST("/", weather)
		v1.GET("/:source_id", source)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(fmt.Sprintf("%s:%d", server, port))
}
