package gmail

import (
	"context"
	"fmt"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/client"
	domain "github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace/gmail"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/logger"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/tracer"
	"github.com/google/wire"
	"github.com/playwright-community/playwright-go"
)

var _ domain.Safety = (*Safety)(nil)

var NewSafetySet = wire.NewSet(
	wire.Bind(new(domain.Safety), new(*Safety)),
	NewSafety,
)

type Safety struct {
	logger logger.Logger
	tracer tracer.Tracer
	web    client.Web
}

func NewSafety(
	logger logger.Logger,
	tracer tracer.Tracer,
	web client.Web,
) (*Safety, error) {
	return &Safety{
		logger: logger,
		tracer: tracer,
		web:    web,
	}, nil
}

const (
	url = "https://admin.google.com/ac/apps/gmail/safety?hl=en"
)

func (c *Safety) Do(ctx context.Context, page playwright.Page) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Gmail.Safety")
	defer span.End()

	if _, err := page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		return fmt.Errorf("failed to visit URL(%s): %w", url, err)
	}

	if err := c.encryptedAttachmentsFromUntrustedSenders(ctx, page); err != nil {
		return fmt.Errorf("failed to check encryptedAttachmentsFromUntrustedSenders: %w", err)
	}

	return nil
}

func (c *Safety) encryptedAttachmentsFromUntrustedSenders(ctx context.Context, page playwright.Page) error {
	text := "Protect against encrypted attachments from untrusted senders"

	enabled, err := c.enabled(ctx, page, text)
	if err != nil {
		return fmt.Errorf("failed to detect encrypted attachments from untrusted senders has enabled: %w", err)
	}

	c.logger.Debug(ctx, "encryptedAttachmentsFromUntrustedSenders", c.logger.Field("enalbed", enabled))

	return nil
}

func (c *Safety) enabled(ctx context.Context, page playwright.Page, text string) (bool, error) {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Gmail.Safety.enabled")
	defer span.End()

	textLocator, err := page.Locator("label[role=presentation]", playwright.PageLocatorOptions{
		HasText: text,
	})
	if err != nil {
		return false, fmt.Errorf("failed to locate element whoch has text(%s): %w", text, err)
	}

	checkboxLocator, err := textLocator.Locator("div[role=checkbox]")
	if err != nil {
		return false, fmt.Errorf("failed to locate element which has checkbox: %w", err)
	}

	checked, err := checkboxLocator.IsChecked()
	if err != nil {
		return false, fmt.Errorf("failed to detect that checkbox is checked: %w", err)
	}

	return checked, nil
}
