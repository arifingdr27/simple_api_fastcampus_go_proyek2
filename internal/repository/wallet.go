package repository

import (
	"context"

	"ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type WalletRepo struct {
	DB *gorm.DB
}

func (r *WalletRepo) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	if err := r.DB.Create(wallet).Error; err != nil {
		return err
	}
	return nil
}
