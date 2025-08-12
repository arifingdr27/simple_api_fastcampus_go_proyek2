package services

import (
	"context"
	"time"

	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	user, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return models.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.LoginResponse{}, err
	}

	token, err := helpers.GenerateToken(ctx, int64(user.ID), user.Username, user.FullName, "jwt", time.Now())
	if err != nil {
		return models.LoginResponse{}, err
	}

	refreshToken, err := helpers.GenerateToken(ctx, int64(user.ID), user.Username, user.FullName, "refresh_token", time.Now())
	if err != nil {
		return models.LoginResponse{}, err
	}

	userSession := models.UserSession{
		UserID:              user.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        time.Now().Add(15 * time.Minute),
		RefreshTokenExpired: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.UserRepo.InsertNewUserSession(ctx, userSession); err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		UserID:       user.ID,
		FullName:     user.FullName,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
		Username:     user.Username,
	}, nil
}
