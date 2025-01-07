package cmd

import (
	// "net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	dotEnv()

	e := echo.New()
	
	// mysql.DataBaseinit()
	// database.RunMigration()

	e.Validator = validate.New()
	// log.Init()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))
	//route end point init
	// router.RouteInt(e.Group("/api/v1"))

	PORT := os.Getenv("APP_PORT")

	// default port 500
	if PORT == "" {
		PORT = "5000"
	}

	e.Logger.Fatal(e.Start(":" + PORT))
}

func dotEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}
