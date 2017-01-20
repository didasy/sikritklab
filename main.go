package main

import (
	"os"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/custommiddleware"
	"github.com/JesusIslam/sikritklab/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(custommiddleware.DeleteOldThread)

	e.GET("/thread/search", handler.ThreadSearch)
	e.GET("/thread/random", handler.ThreadRandom)
	e.GET("/thread/id/:id", handler.ThreadGetByID)
	e.POST("/thread/new", handler.ThreadNew)
	e.POST("/thread/id/:id", handler.ThreadReplyByID)

	host := ":8080"
	if os.Getenv("SIKRIT_HOST") != "" {
		host = os.Getenv("SIKRIT_HOST")
	}
	lowger.Fatal(e.Start(host))
}
