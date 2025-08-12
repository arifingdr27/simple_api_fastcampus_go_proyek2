package repository

import (
	"context"
	"fmt"

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
		fmt.Println("err")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("121")
	fmt.Println(user)
	if user.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, session models.UserSession) error {
	return r.DB.Create(&session).Error
}
