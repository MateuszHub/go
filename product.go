package main

import (
	"net/http"
    "github.com/labstack/echo/v4"  
	"strconv"	
    "gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func getProduct(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	DB.First(&product, id)
	return c.JSON(http.StatusOK, product)
}


func getAllProducts(c echo.Context) error {
	var products []Product
	DB.Find(&products)
	return c.JSON(http.StatusOK, products)
}


func addProduct(c echo.Context) error {
	p := &Product{}	
	if err := c.Bind(p); err != nil {
		return err
	}
	DB.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func RegisterProductApi(echo *echo.Echo)  {	
	DB.AutoMigrate(&Product{})
	echo.GET("/products", getAllProducts)
    echo.GET("/products/:id", getProduct)
    echo.POST("/products", addProduct)
}