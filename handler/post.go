package handler

import (
	"log"
	"net/http"

	"github.com/gregdel/pushover"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SendGenericToPushover(c echo.Context) (err error) {

	app := pushover.New(h.Config.PushoverApiKey)

	recipient := pushover.NewRecipient(h.Config.PushoverUserKey)

	message := pushover.NewMessage("Hello !")

	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	return c.JSON(http.StatusOK, `{"message": "Message was sent. Response: `+response.String()+`"}`)
}
