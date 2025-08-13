package services

import (
	"context"
	"time"

	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaims helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	token, err := helpers.GenerateToken(ctx, int64(tokenClaims.UserID), tokenClaims.Email, tokenClaims.Username, tokenClaims.Fullname, "token", time.Now())
	if err != nil {
		return models.RefreshTokenResponse{}, err
	}

	err = s.UserRepo.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return models.RefreshTokenResponse{}, err
	}

	return models.RefreshTokenResponse{
		Token: token,
	}, nil
}
