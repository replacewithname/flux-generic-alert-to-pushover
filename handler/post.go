package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gregdel/pushover"
	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	InvolvedObject struct {
		Kind      string `json:"kind"`
		Namespace string `json:"namespace"`
	} `json:"involvedObject"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

func (h *Handler) SendGenericToPushover(c echo.Context) (err error) {

	// 1. Gather Data and check validity
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, `{"message": "Request body could not be read"}`)
	}

	body := &RequestBody{}

	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, `{"message": "Request body could not be parsed"}`)
	}

	// 2. Prepare Pushover and forward message
	app := pushover.New(h.Config.PushoverApiKey)

	recipient := pushover.NewRecipient(h.Config.PushoverUserKey)

	message := pushover.NewMessageWithTitle(body.Message, body.Reason)

	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	return c.JSON(http.StatusOK, `{"message": "Message was sent. Response: `+response.String()+`"}`)
}
