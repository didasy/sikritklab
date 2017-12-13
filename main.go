package main

import (
	"os"
	"strings"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/JesusIslam/sikritklab/internal/custommiddleware"
	"github.com/JesusIslam/sikritklab/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	if os.Getenv(constant.EnvGinMode) != constant.EnvGinModeValueRelease {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv(constant.EnvOrigins), ","), // * or your url (https://example.com)
		AllowMethods:     strings.Split(os.Getenv(constant.EnvMethods), ","), // should be GET,POST
		AllowCredentials: true,
	}))
	r.Use(custommiddleware.DeleteOldThread())

	r.GET("/thread/search", handler.ThreadSearch)
	r.GET("/thread/random", handler.ThreadRandom)
	r.GET("/thread/id/:id", handler.ThreadGetByID)
	r.POST("/thread/new", custommiddleware.CheckCaptcha(), handler.ThreadNew)
	r.POST("/thread/id/:id", custommiddleware.CheckCaptcha(), handler.ThreadReplyByID)

	host := constant.DefaultHost
	lowger.Info(constant.InfoListeningHost, host)
	if os.Getenv(constant.EnvHost) != "" {
		host = os.Getenv(constant.EnvHost)
	}
	lowger.Fatal(r.Run(host))
}
