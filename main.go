package main

import (
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dbVar := os.Getenv("NUC_DB")

	db, err := gorm.Open(mysql.Open(dbVar), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	if err != nil {
		panic("failed to connect database")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
