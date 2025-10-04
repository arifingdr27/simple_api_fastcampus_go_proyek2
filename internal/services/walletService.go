package services

import (
	"context"

	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
)

type WalletService struct {
	WalletRepo interfaces.IWalletRepo
}

func (s *WalletService) Create(ctx context.Context, wallet *models.Wallet) error {
	return s.WalletRepo.CreateWallet(ctx, wallet)
}
