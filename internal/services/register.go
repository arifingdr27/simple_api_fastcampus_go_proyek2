package services

import (
	"context"
	"time"

	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo interfaces.IRegisterRepository
}

func (rs RegisterService) Register(ctx context.Context, register models.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	register.Password = string(hashPassword)
	register.CreatedAt = time.Now()
	register.UpdatedAt = time.Now()
	if err := rs.RegisterRepo.InsertNewUser(ctx, register); err != nil {
		return nil, err
	}
	return register, nil
}
