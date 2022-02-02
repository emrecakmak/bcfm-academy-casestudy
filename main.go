package main

import (
	"net/http"

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

func main() {
	e := echo.New()

	e.GET("/", userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
