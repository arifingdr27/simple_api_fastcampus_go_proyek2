package memberships

import (
	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(user *memberships.User) error
	GetUser(email, username string, id int) (*memberships.User, error)
}

type service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}
