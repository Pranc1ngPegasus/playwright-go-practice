package googleworkspace

import (
	"context"
)

type GoogleWorkspace interface {
	Do(context.Context, LoginInput) error
}

type (
	LoginInput struct {
		Email    string
		Password string
		TOTP     string
	}
)
