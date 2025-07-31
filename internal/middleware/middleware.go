package middleware

import (
	"errors"
	"net/http"
	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJwt
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
		}
		userId, username, err := jwt.ValidationToken(header, secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		ctx.Set("userID", userId)
		ctx.Set("username", username)
		ctx.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJwt
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
		}
		userId, username, err := jwt.ValidationTokenWithoutExpired(header, secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		ctx.Set("userID", userId)
		ctx.Set("username", username)
		ctx.Next()
	}
}
