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
	UpdateTokenByRefreshToken(ctx context.Context, token string, refresh_token string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error)
}

type IUserService interface {
	Register(ctx context.Context, register models.User) (interface{}, error)
}
