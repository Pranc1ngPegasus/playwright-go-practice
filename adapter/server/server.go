package server

import (
	"net/http"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/configuration"
)

func NewServer(
	config configuration.Configuration,
	handler http.Handler,
) *http.Server {
	return &http.Server{
		Addr:    ":" + config.Server().Port,
		Handler: handler,
	}
}
