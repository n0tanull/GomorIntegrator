package service

import (
	"context"
	"simpleauth/internal/repo/client"
	"simpleauth/internal/repo/user"
	"simpleauth/pkg/melog"
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
	mel      *melog.Logger
}

func NewService(mel *melog.Logger) Service {
	return &service{
		repo:     user.NewRepo(mel),
		provider: client.NewSocialProvider(),
		mel:      mel,
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
