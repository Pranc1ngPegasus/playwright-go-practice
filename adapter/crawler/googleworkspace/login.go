package googleworkspace

import (
	"context"
	"fmt"
	"strings"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/client"
	domain "github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/logger"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/tracer"
	"github.com/google/wire"
	"github.com/playwright-community/playwright-go"
)

var _ domain.Login = (*Login)(nil)

var NewLoginSet = wire.NewSet(
	wire.Bind(new(domain.Login), new(*Login)),
	NewLogin,
)

type Login struct {
	logger logger.Logger
	tracer tracer.Tracer
	web    client.Web
}

func NewLogin(
	logger logger.Logger,
	tracer tracer.Tracer,
	web client.Web,
) (*Login, error) {
	return &Login{
		logger: logger,
		tracer: tracer,
		web:    web,
	}, nil
}

const (
	url = "https://admin.google.com/AdminHome?hl=en"
)

func (c *Login) Do(ctx context.Context, page playwright.Page, input domain.LoginInput) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Login")
	defer span.End()

	if _, err := page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		return fmt.Errorf("failed to visit URL(%s): %w", url, err)
	}

	// メールアドレスを入力
	if err := c.fillEmail(ctx, page, input.Email); err != nil {
		return fmt.Errorf("failed to fill email: %w", err)
	}

	// パスワードを入力
	if err := c.fillPassword(ctx, page, input.Password); err != nil {
		return fmt.Errorf("failed to fill email: %w", err)
	}

	// MFA認証
	if err := c.challengeMFA(ctx, page, input.TOTP); err != nil {
		return fmt.Errorf("failed to challenge MFA: %w", err)
	}

	return nil
}

func (c *Login) fillEmail(ctx context.Context, page playwright.Page, email string) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Login.fillEmail")
	defer span.End()

	emailInput, err := page.Locator("input[type=email]")
	if err != nil {
		return fmt.Errorf("failed to locate email input: %w", err)
	}

	if err := emailInput.Fill(email); err != nil {
		return fmt.Errorf("failed to fill email input: %w", err)
	}

	nextButton, err := page.Locator("#identifierNext")
	if err != nil {
		return fmt.Errorf("failed to locate next button: %w", err)
	}

	if err := nextButton.Click(); err != nil {
		return fmt.Errorf("failed to click next button: %w", err)
	}

	if _, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		return fmt.Errorf("failed to wait navigation: %w", err)
	}

	return nil
}

func (c *Login) fillPassword(ctx context.Context, page playwright.Page, password string) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Login.fillPassword")
	defer span.End()

	passwordInput, err := page.Locator("input[type=password]")
	if err != nil {
		return fmt.Errorf("failed to locate password input: %w", err)
	}

	if err := passwordInput.Fill(password); err != nil {
		return fmt.Errorf("failed to fill password input: %w", err)
	}

	nextButton, err := page.Locator("#passwordNext")
	if err != nil {
		return fmt.Errorf("failed to locate next button: %w", err)
	}

	if err := nextButton.Click(); err != nil {
		return fmt.Errorf("failed to click next button: %w", err)
	}

	if _, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		return fmt.Errorf("failed to wait navigation: %w", err)
	}

	return nil
}

func (c *Login) challengeMFA(ctx context.Context, page playwright.Page, totp string) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.Login.checkTOTP")
	defer span.End()

	switch {
	case strings.Contains(page.URL(), "https://admin.google.com/AdminHome."):
		c.logger.Debug(ctx, "MFA skipped")

		return nil
	case strings.Contains(page.URL(), "challenge/selection"):
		totpSelection, err := page.Locator("div[role=link][data-challengetype=\"6\"]")
		if err != nil {
			return fmt.Errorf("failed to locate totp selection: %w", err)
		}

		if err := totpSelection.Click(); err != nil {
			return fmt.Errorf("failed to click totp selection: %w", err)
		}

		if _, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		}); err != nil {
			return fmt.Errorf("failed to wait navigation: %w", err)
		}

		fallthrough
	case strings.Contains(page.URL(), "challenge/totp"):
		totpInput, err := page.Locator("input[type=tel]")
		if err != nil {
			return fmt.Errorf("failed to locate TOTP input: %w", err)
		}

		if err := totpInput.Fill(totp); err != nil {
			return fmt.Errorf("failed to fill TOTP input: %w", err)
		}

		nextButton, err := page.Locator("#totpNext")
		if err != nil {
			return fmt.Errorf("failed to locate next button: %w", err)
		}

		if err := nextButton.Click(); err != nil {
			return fmt.Errorf("failed to click next button: %w", err)
		}

		if _, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		}); err != nil {
			return fmt.Errorf("failed to wait navigation: %w", err)
		}

		fallthrough
	default:
		if _, err := page.WaitForNavigation(playwright.PageWaitForNavigationOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		}); err != nil {
			return fmt.Errorf("failed to wait navigation: %w", err)
		}

		return nil
	}
}
