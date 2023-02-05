package googleworkspace

import (
	"context"

	"github.com/playwright-community/playwright-go"
)

type Login interface {
	Do(context.Context, playwright.Page, LoginInput) error
}
