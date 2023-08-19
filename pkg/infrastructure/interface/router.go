package pkg

import (
	"context"

	auth "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/authentication"
)

type Router interface {
	Connect(ctx context.Context, asset auth.Asset) error
	Login(ctx context.Context, user auth.User) error
	ShowRouter(ctx context.Context) error
	Backup(ctx context.Context) (string, error)
}
