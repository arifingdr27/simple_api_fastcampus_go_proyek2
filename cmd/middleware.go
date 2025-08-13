package cmd

import (
	"log"
	"net/http"
	"time"

	"ewallet-ums/helpers"

	"github.com/gin-gonic/gin"
)

func MiddlewareValidateAuth(ctx *gin.Context, depedency Depedency) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
	}

	// _, err := depe.GetUserSessionByToken(ctx.Request.Context(), auth)
	_, err := depedency.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
	}

	ctx.Set("token", claim)

	ctx.Next()
}
