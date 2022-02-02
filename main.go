package main

import (
	"fmt"
	"net/http"

	owm "github.com/briandowns/openweathermap"
	"github.com/labstack/echo/v4"
)

func userHandler(c echo.Context) error {
	var myInfo struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}
	myInfo.Name = "Emre"
	myInfo.Surname = "Çakmak"

	return c.JSON(http.StatusOK, &myInfo)
}

func tempatureHandler(c echo.Context) error {
	cityName := c.QueryParam("city")

	type weatherData struct {
		Tempature float64 `json:"tempature"`
	}
	//openweather api connection
	w, err := owm.NewCurrent("C", "tr", "0b1bfdc3222058877e875f5bb0259ddf")
	if err != nil {
		return fmt.Errorf("API Error")
	}

	w.CurrentByName(cityName)

	weatherTemp := &weatherData{
		Tempature: w.Main.Temp,
	}

	return c.JSON(http.StatusOK, &weatherTemp)
}

func main() {
	e := echo.New()

	e.GET("/", userHandler)
	e.GET("/tempature", tempatureHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
