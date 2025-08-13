package repository

import (
	"context"

	"ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user models.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session models.UserSession) error {
	return r.DB.Create(&session).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_session WHERE token = ?", token).Error
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, token string, refresh_token string) error {
	return r.DB.Exec("UPDATE user_session SET token = ? WHERE refresh_token = ?", token, refresh_token).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var session *models.UserSession
	if err := r.DB.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, err
	}
	if session.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return session, nil
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error) {
	var session *models.UserSession
	if err := r.DB.Where("refresh_token = ?", refreshToken).First(&session).Error; err != nil {
		return nil, err
	}
	if session.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return session, nil
}
