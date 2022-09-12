package main

import (
	"net/http"
    "github.com/labstack/echo/v4"
	"strconv"	
    "gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Telephone  string
	Mail string
}


func getInfo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
	var info Info
	DB.First(&info, id)
	return c.JSON(http.StatusOK, info)
}


func getAllInfo(c echo.Context) error {
	var infos []Info
	DB.Find(&infos)
	return c.JSON(http.StatusOK, infos)
}


func addInfo(c echo.Context) error {
	p := &Info{}	
	if err := c.Bind(p); err != nil {
		return err
	}
	DB.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func RegisterInfoApi(echo *echo.Echo)  {
	DB.AutoMigrate(&Info{})

	echo.GET("/info", getAllInfo)
    echo.GET("/info/:id", getInfo)
    echo.POST("/info", addInfo)
}