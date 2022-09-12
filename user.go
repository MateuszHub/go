package main

import (
	"net/http"
    "github.com/labstack/echo/v4"
	"strconv"	
    "gorm.io/gorm"
	"fmt"
)

type User struct {
	gorm.Model
	Name  string
	Surname string
	Orders []Order
}


func getUser(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
	var user User
	DB.Model(&User{}).Preload("Orders").Preload("Orders.Products").First(&user, id)
	
    fmt.Println("Hello, World!")
	fmt.Println(user.ID, user.Orders)
	return c.JSON(http.StatusOK, user.Orders)
}


func getAllUser(c echo.Context) error {
	var users []User
	DB.Preload("Orders").Preload("Orders.Products").Find(&users)
	return c.JSON(http.StatusOK, users)
}


func addUser(c echo.Context) error {
	p := &User{}	
	if err := c.Bind(p); err != nil {
		return err
	}
	DB.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func RegisterUserApi(echo *echo.Echo)  {
	DB.AutoMigrate(&User{})

	echo.GET("/users", getAllUser)
    echo.GET("/users/:id", getUser)
    echo.POST("/users", addUser)
}