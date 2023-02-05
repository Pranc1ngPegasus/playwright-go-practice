package googleworkspace

import (
	"context"

	"github.com/playwright-community/playwright-go"
)

type GmailSafety interface {
	Do(ctx context.Context, page playwright.Page) error
}
