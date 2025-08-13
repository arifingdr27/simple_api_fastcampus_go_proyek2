package interfaces

import (
	"context"

	"ewallet-ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InsertNewUserSession(ctx context.Context, userSession models.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
}

type IUserService interface {
	Register(ctx context.Context, register models.User) (interface{}, error)
}
