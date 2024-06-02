package user

import "context"

type Repo interface {
	AddSocial(telegram string, ctx context.Context) (string, error)
	AddUser(email, password string, ctx context.Context) (string, error)
	GetUser(id string, ctx context.Context) (string, error)
	GetUsers(ctx context.Context) ([]string, error)
}
