package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (h *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	token, ok := c.Get("token")
	if !ok {
		helpers.SendResponseHttp(c, 400, "invalid token", nil)
		return
	}

	tokenClaims, ok := token.(*helpers.ClaimToken)
	if !ok {
		helpers.SendResponseHttp(c, 400, "invalid token2", nil)
		return
	}

	resp, err := h.RefreshTokenService.RefreshToken(c, refreshToken, *tokenClaims)
	if err != nil {
		helpers.SendResponseHttp(c, 500, err.Error(), nil)

		return
	}

	helpers.SendResponseHttp(c, 200, "token refreshed successfully", resp)
}
