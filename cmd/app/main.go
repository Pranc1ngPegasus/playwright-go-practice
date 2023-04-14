package main

import (
	"context"
	"os"
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
	if err := app.server.ListenAndServe(); err != nil {
		app.logger.Error(ctx, "failed to start server", app.logger.Field("err", err))
	}

	if err := app.tracer.Shutdown(ctx); err != nil {
		app.logger.Error(ctx, "failed to stop tracer", app.logger.Field("err", err))

		return
	}
}
