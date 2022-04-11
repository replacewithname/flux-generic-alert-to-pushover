package main

import (
	"flux-generic-alert-to-pushover/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/send-generic-to-pushover", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	h := &handler.Handler{}

	if assert.NoError(t, h.SendGenericToPushover(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, "test", rec.Body.String())
	}
}
