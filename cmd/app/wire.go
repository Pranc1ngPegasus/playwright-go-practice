//go:build wireinject
// +build wireinject

package main

import (
	gwCrawler "github.com/Pranc1ngPegasus/playwright-go-practice/adapter/crawler/googleworkspace"
	gmailCrawler "github.com/Pranc1ngPegasus/playwright-go-practice/adapter/crawler/googleworkspace/gmail"
	domaincrawler "github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace"
	domainlogger "github.com/Pranc1ngPegasus/playwright-go-practice/domain/logger"
	domaintracer "github.com/Pranc1ngPegasus/playwright-go-practice/domain/tracer"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/client"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/configuration"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/logger"
	"github.com/Pranc1ngPegasus/playwright-go-practice/infra/tracer"
	"github.com/google/wire"
)

type app struct {
	logger  domainlogger.Logger
	tracer  domaintracer.Tracer
	crawler domaincrawler.GoogleWorkspace
}

func initialize() (*app, error) {
	wire.Build(
		configuration.NewConfigurationSet,
		logger.NewLoggerSet,
		tracer.NewTracerSet,
		gwCrawler.NewLoginSet,
		gwCrawler.NewGoogleWorkspaceSet,
		client.NewWebSet,
		gmailCrawler.NewSafetySet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
