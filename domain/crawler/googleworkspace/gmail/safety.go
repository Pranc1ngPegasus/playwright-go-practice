package gmail

import (
	"context"

	"github.com/playwright-community/playwright-go"
)

type Safety interface {
	Do(ctx context.Context, page playwright.Page) error
}
