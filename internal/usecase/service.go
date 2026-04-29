package usecase

import "context"

type Service interface {
	ListAll(ctx context.Context) error
}

type svc struct {
	//repo nil,
}

func NewService() Service {
	return nil //&svc{}
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}
