package service

import (
	"context"
	"simpleauth/internal/repo/client"
	"simpleauth/internal/repo/user"
)

type Service interface {
	AddUser(email, password string, ctx context.Context) (string, error)
	AddSocial(telegram string, ctx context.Context) (string, error)
	GetUser(id string, ctx context.Context) (string, error)
	GetUsers(ctx context.Context) ([]string, error)
}

type service struct {
	repo     user.Repo
	provider client.SocialNetworkProvider
}

func NewService(repo user.Repo) Service {
	return &service{
		repo:     repo,
		provider: client.NewSocialProvider(),
	}
}

func (s *service) AddUser(email, password string, ctx context.Context) (string, error) {
	return s.repo.AddUser(email, password, ctx)
}

func (s *service) AddSocial(telegram string, ctx context.Context) (string, error) {
	return s.repo.AddSocial(telegram, ctx)

}

func (s *service) GetUser(id string, ctx context.Context) (string, error) {
	return s.repo.GetUser(id, ctx)
}

func (s *service) GetUsers(ctx context.Context) ([]string, error) {
	return s.repo.GetUsers(ctx)
}
