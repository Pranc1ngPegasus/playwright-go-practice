package handler

import (
	"fmt"
	"net/http"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/configuration"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace"
	"github.com/google/wire"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	handler http.Handler
	config  configuration.Configuration
	crawler googleworkspace.GoogleWorkspace
}

func NewHandler(
	config configuration.Configuration,
	crawler googleworkspace.GoogleWorkspace,
) *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		handler: mux,
		config:  config,
		crawler: crawler,
	}

	mux.HandleFunc("/healthcheck", h.healthcheck)
	mux.HandleFunc("/scan", h.scan)

	return h
}

func (h *Handler) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func (h *Handler) scan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.crawler.Do(ctx, googleworkspace.LoginInput{
		Email:    h.config.Scan().Email,
		Password: h.config.Scan().Password,
		TOTP:     h.config.Scan().TOTP,
	}); err != nil {
		fmt.Fprintf(w, "failed to crawl: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
