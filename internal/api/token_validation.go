package api

import (
	"context"
	"fmt"

	"ewallet-ums/cmd/proto/tokenvalidation"
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.Token
		log   = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is empty")
		log.Error(err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constants.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
			Email:    claimToken.Email,
		},
	}, nil
}

// rpc ValidateToken (TokenRequest) returns (TokenResponse);

func (s *TokenValidationHandler) ValidateDataUser(ctx context.Context, token *tokenvalidation.DataUser) (*tokenvalidation.UserResponse, error) {
	if token.Token == "" {
		err := fmt.Errorf("token is empty")
		return nil, err
	}

	// claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	// if err != nil {
	// 	return nil, err
	// }

	// userData := map[string]interface{}{
	// 	"user_id":   1,
	// 	"username":  "testuser",
	// 	"full_name": "Test User",
	// 	"email":     "testuser@example.com",
	// }

	return &tokenvalidation.UserResponse{
		Message: "success",
		Data: &tokenvalidation.UserAJA{
			Username: "testuser",
			FullName: "Test User",
		},
	}, nil
}
