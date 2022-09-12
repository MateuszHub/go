package main
import (
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
)

var dsn = "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})


func main() {

	e := echo.New()
	RegisterProductApi(e)
	RegisterInfoApi(e)
	RegisterOrderApi(e)
	RegisterUserApi(e)
	RegisterPlaceApi(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"hello world")
	})
	e.Logger.Fatal(e.Start(":3000"))
}