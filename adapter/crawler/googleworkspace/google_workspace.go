package googleworkspace

import (
	"context"
	"fmt"

	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/client"
	domain "github.com/Pranc1ngPegasus/playwright-go-practice/domain/crawler/googleworkspace"
	"github.com/Pranc1ngPegasus/playwright-go-practice/domain/tracer"
	"github.com/google/wire"
	"github.com/playwright-community/playwright-go"
	"github.com/volatiletech/null/v8"
)

var _ domain.GoogleWorkspace = (*GoogleWorkspace)(nil)

var NewGoogleWorkspaceSet = wire.NewSet(
	wire.Bind(new(domain.GoogleWorkspace), new(*GoogleWorkspace)),
	NewGoogleWorkspace,
)

type GoogleWorkspace struct {
	tracer tracer.Tracer
	web    client.Web
	login  domain.Login
	safety domain.GmailSafety
}

func NewGoogleWorkspace(
	tracer tracer.Tracer,
	web client.Web,
	login domain.Login,
	safety domain.GmailSafety,
) *GoogleWorkspace {
	return &GoogleWorkspace{
		tracer: tracer,
		web:    web,
		login:  login,
		safety: safety,
	}
}

func (c *GoogleWorkspace) Do(ctx context.Context, input domain.LoginInput) error {
	ctx, span := c.tracer.Tracer().Start(ctx, "crawler.GoogleWorkspace.GoogleWorkspace")
	defer span.End()

	browserCtx, err := c.web.NewContext()
	if err != nil {
		return fmt.Errorf("failed to initialize crawler: %w", err)
	}

	if err := browserCtx.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: null.BoolFrom(true).Ptr(),
		Snapshots:   null.BoolFrom(true).Ptr(),
	}); err != nil {
		return fmt.Errorf("failed to start tracer: %w", err)
	}

	page, err := browserCtx.NewPage()
	if err != nil {
		return fmt.Errorf("failed to create new page: %w", err)
	}

	if err := c.login.Do(ctx, page, input); err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}

	otherPage, err := browserCtx.NewPage()
	if err != nil {
		return fmt.Errorf("failed to create new page: %w", err)
	}

	if err := c.safety.Do(ctx, otherPage); err != nil {
		return fmt.Errorf("failed to crawl Gmail safety: %w", err)
	}

	if err := browserCtx.Tracing().Stop(playwright.TracingStopOptions{
		Path: null.StringFrom("trace.zip").Ptr(),
	}); err != nil {
		return fmt.Errorf("failed to start tracer: %w", err)
	}

	return nil
}
