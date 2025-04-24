package service

import "github.com/PeterKWIlliams/my-to-do-go/internal/config"

type Service struct {
	Config *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		Config: cfg,
	}
}
