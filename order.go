package main

import (
	"net/http"
    "github.com/labstack/echo/v4"
	"strconv"	
    "gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Products []Product `gorm:"many2many:order_product;"`
	UserID uint
}

type OrderSimple struct {
	ProductId []uint
	UserId uint
}

func getOrder(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
	var order Order
	DB.Model(&Order{}).Preload("Products").First(&order, id)
	return c.JSON(http.StatusOK, order)
}


func getAllOrder(c echo.Context) error {
	var orders []Order
	DB.Preload("Products").Find(&orders)
	return c.JSON(http.StatusOK, orders)
}


func addOrder(c echo.Context) error {
	tmp := &OrderSimple{}
	if err := c.Bind(tmp); err != nil {
		return err
	}
	var products []Product
	DB.Where("ID IN ?", tmp.ProductId).Find(&products)
	p := Order{Products: products, UserID: tmp.UserId}
	var user User
	DB.First(&user, tmp.UserId)
	DB.Model(&user).Association("Orders").Append(&p)
	DB.Model(&p).Association("Products").Append(p.Products)
	return c.JSON(http.StatusCreated, p)
}

func RegisterOrderApi(echo *echo.Echo)  {
	DB.AutoMigrate(&Order{})

	echo.GET("/order", getAllOrder)
    echo.GET("/order/:id", getOrder)
    echo.POST("/order", addOrder)
}