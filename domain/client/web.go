package client

import (
	playwright "github.com/playwright-community/playwright-go"
)

type Web interface {
	NewContext() (playwright.BrowserContext, error)
}
