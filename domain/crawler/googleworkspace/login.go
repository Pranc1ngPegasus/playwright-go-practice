package googleworkspace

import "context"

type Login interface {
	Do(context.Context, LoginInput) error
}

type (
	LoginInput struct {
		Email    string
		Password string
		TOTP     string
	}
)
