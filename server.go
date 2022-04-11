package main

import (
	"flux-generic-alert-to-pushover/config"
	"flux-generic-alert-to-pushover/handler"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := &handler.Handler{Config: &config}

	e.POST("/send-generic-to-pushover", h.SendGenericToPushover)

	e.Logger.Fatal(e.Start(":1323"))
}
