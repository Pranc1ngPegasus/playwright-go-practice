package main

import (
	"context"
	"os"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace"
)

var (
	email    string = ""
	password string = ""
	totp     string = ""
)

func init() {
	email = os.Getenv("GOOGLE_WORKSPACE_EMAIL")
	password = os.Getenv("GOOGLE_WORKSPACE_PASSWORD")
	totp = os.Getenv("GOOGLE_WORKSPACE_TOTP")
}

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := app.crawler.Do(ctx, googleworkspace.LoginInput{
		Email:    email,
		Password: password,
		TOTP:     totp,
	}); err != nil {
		app.logger.Error(ctx, "failed to crawl", app.logger.Field("err", err))

		return
	}

	if err := app.tracer.Shutdown(ctx); err != nil {
		app.logger.Error(ctx, "failed to stop tracer", app.logger.Field("err", err))

		return
	}
}
