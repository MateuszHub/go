package main

import (
	"net/http"
    "github.com/labstack/echo/v4"
	"strconv"	
    "gorm.io/gorm"
)

type Place struct {
	gorm.Model
	street string
	city string
}


func getPlace(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
	var place Place
	DB.First(&place, id)
	return c.JSON(http.StatusOK, place)
}


func getAllPlace(c echo.Context) error {
	var places []Place
	DB.Find(&places)
	return c.JSON(http.StatusOK, places)
}


func addPlace(c echo.Context) error {
	p := &Place{}	
	if err := c.Bind(p); err != nil {
		return err
	}
	DB.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func RegisterPlaceApi(echo *echo.Echo)  {
	DB.AutoMigrate(&Place{})

	echo.GET("/place", getAllPlace)
    echo.GET("/place/:id", getPlace)
    echo.POST("/place", addPlace)
}