package main

import (
	"myapp/handlers"
	"myapp/models"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbVar := os.Getenv("NUC_DB")
	e := echo.New()

	db, err := gorm.Open(postgres.Open(dbVar), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	dbHandler := handlers.NewDBHandler(db)

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", hello)
	e.POST("/users", dbHandler.CreateUser)
	e.GET("/users/:id", dbHandler.GetUser)
	e.GET("/users", dbHandler.GetAllUsers)
	e.PUT("/users/:id", dbHandler.UpdateUser)
	e.DELETE("/users/:id", dbHandler.DeleteUser)
	e.Logger.Fatal(e.Start(":3000"))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
}
