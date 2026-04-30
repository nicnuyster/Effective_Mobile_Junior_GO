package usecase

import (
	"context"

	repo "github.com/nicnuyster/Effective_Mobile_Junior_GO/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListAll(ctx context.Context) ([]repo.Subscribtions, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListAll(ctx context.Context) ([]repo.Subscribtions, error) {
	return s.repo.ListAll(ctx)
}

func (s *svc) ListOne(ctx context.Context) ([]repo.Subscribtions, error) {
	return s.repo.ListAll(ctx)
}
