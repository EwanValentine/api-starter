package main

import (
	"bitbucket.org/ewanvalentine/api-starter/drivers"
	"bitbucket.org/ewanvalentine/api-starter/handlers"
	"bitbucket.org/ewanvalentine/api-starter/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"log"
	"runtime"
)

func Init() {

	// Verbose logging
	log.SetFlags(log.Lshortfile)

	// Use all available cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	Init()

	datastore := drivers.DB()

	// Migrate changes
	drivers.DB().AutoMigrate(&models.Thing{})

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.POST,
			echo.DELETE,
			echo.PATCH,
			echo.OPTIONS,
		},
	}))

	repository := models.NewRepository(datastore.C("content"))
	handlers := handlers.NewHandler(repository)

	e.GET("/api/v1/things", handlers.FindAll)
	e.POST("/api/v1/things", handlers.Insert)
	e.GET("/api/v1/things/:id", handlers.Find)
	e.PATCH("/api/v1/things/:id", handlers.Update)
	e.DELETE("/api/v1/things/:id", handlers.Remove)

	e.Run(fasthttp.New(":1000"))
}
