package handler

import "flux-generic-alert-to-pushover/config"

type (
	Handler struct {
		Config *config.Config
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
