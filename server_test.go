package main

import (
	"flux-generic-alert-to-pushover/config"
	"flux-generic-alert-to-pushover/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gregdel/pushover"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	conf := &config.Config{PushoverApiKey: "a34wsc2e7xdwbbwntgewqybys4fmn9", PushoverUserKey: "uff7qpdvsa9czkjetx4m7i7c9ff689"}

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/send-generic-to-pushover", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// For some reason, this is the only thing which is checked by the framework
		w.Header().Add("x-limit-app-limit", "10000")
		w.Header().Add("x-limit-app-remaining", "9994")
		w.Header().Add("x-limit-app-reset", "1651381200")

		w.Write([]byte(`{"status":1,"request":"bf59e63c-8519-401e-ad74-74b777bf9b85"}`))
	}))

	pushover.APIEndpoint = server.URL

	h := &handler.Handler{Config: conf}

	if assert.NoError(t, h.SendGenericToPushover(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, "test", rec.Body.String())
	}
}
