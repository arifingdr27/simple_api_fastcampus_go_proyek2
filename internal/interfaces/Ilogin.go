package interfaces

import (
	"context"

	"ewallet-ums/internal/models"
)

type IloginRepository interface{}

type IloginService interface {
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)
}
