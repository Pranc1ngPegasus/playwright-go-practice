//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	gwCrawler "github.com/Pranc1ngPegasus/playwright-go-practice/adapter/crawler/googleworkspace"
	"github.com/Pranc1ngPegasus/playwright-go-practice/adapter/handler"
	"github.com/Pranc1ngPegasus/playwright-go-practice/adapter/server"
	domainlogger "github.com/Pranc1ngPegasus/playwright-go-practice/domain/logger"
	domaintracer "github.com/Pranc1ngPegasus/playwright-go-practice/domain/tracer"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/client"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/configuration"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/logger"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/tracer"
	"github.com/google/wire"
)

type app struct {
	logger domainlogger.Logger
	server *http.Server
	tracer domaintracer.Tracer
}

func initialize() (*app, error) {
	wire.Build(
		configuration.NewConfigurationSet,
		logger.NewLoggerSet,
		tracer.NewTracerSet,
		handler.NewHandlerSet,
		server.NewServer,
		gwCrawler.NewLoginSet,
		gwCrawler.NewGoogleWorkspaceSet,
		client.NewWebSet,
		gwCrawler.NewGmailSafetySet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
