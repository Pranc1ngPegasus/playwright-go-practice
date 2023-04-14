package client

import (
	"fmt"

	domain "github.com/Pranc1ngPegasus/playwright-go-practice/domain/client"
	"github.com/google/wire"
	playwright "github.com/playwright-community/playwright-go"
)

var _ domain.Web = (*Web)(nil)

var NewWebSet = wire.NewSet(
	wire.Bind(new(domain.Web), new(*Web)),
	NewWeb,
)

type Web struct {
	browser playwright.Browser
}

func NewWeb() (*Web, error) {
	pw, err := playwright.Run(&playwright.RunOptions{
		DriverDirectory:     "/root/.cache",
		SkipInstallBrowsers: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start playwright: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Channel:  null.StringFrom("chrome").Ptr(),
		// Headless: null.BoolFrom(false).Ptr(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to launch browser: %w", err)
	}

	return &Web{
		browser: browser,
	}, nil
}

func (c *Web) NewContext() (playwright.BrowserContext, error) {
	ctx, err := c.browser.NewContext()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize browser context: %w", err)
	}

	return ctx, nil
}
