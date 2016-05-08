package main

import (
	"encoding/json"
	"fmt"
	"github.com/EwanValentine/api-starter/drivers"
	"github.com/EwanValentine/api-starter/handlers"
	"github.com/EwanValentine/api-starter/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
	"runtime"
)

// Init - Bootstrap runtime options
func Init() {

	// Verbose logging
	log.SetFlags(log.Lshortfile)

	// Use all available cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Config - Config object
type Config struct {
	Port   string `json:"port"`
	DBHost string `json:"db_host"`
	DBPass string `json:"db_pass"`
	DBUser string `json:"db_user"`
	DBPort int    `json:"db_port"`
	DBName string `json:"db_name"`
}

func main() {

	Init()

	var config Config

	// Configure
	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&config)

	// Configuration file incorrect or not found
	if err != nil {
		panic(err)
	}

	datastore := drivers.DB(
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBName,
	)

	// Migrate changes
	datastore.AutoMigrate(&models.Thing{})

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

	thingRepository := models.NewThingRepository(datastore)
	handlers := handlers.NewThingHandler(thingRepository)

	e.GET("/api/v1/things", handlers.FindAll)
	e.POST("/api/v1/things", handlers.Insert)
	/*
		e.GET("/api/v1/things/:id", handlers.Find)
		e.PATCH("/api/v1/things/:id", handlers.Update)
		e.DELETE("/api/v1/things/:id", handlers.Remove)
	*/

	fmt.Println("Connecting on port " + config.Port)

	e.Run(fasthttp.New(":" + config.Port))
}
