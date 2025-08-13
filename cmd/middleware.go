package cmd

import (
	"log"
	"net/http"
	"time"

	"ewallet-ums/helpers"

	"github.com/gin-gonic/gin"
)

func (d *Depedency) MiddlewareValidateAuth(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	// _, err := depe.GetUserSessionByToken(ctx.Request.Context(), auth)
	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}

func (d *Depedency) MiddlewareRefreshToken(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}
	_, err := d.UserRepository.GetUserSessionByRefreshToken(ctx, auth)
	if err != nil {
		log.Println("failed to get user session on DB: ", err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
	}

	claim, err := helpers.ValidateToken(ctx, auth)
	if err != nil {
		log.Println(err)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("jwt token is expired: ", claim.ExpiresAt)
		helpers.SendResponseHttp(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)

	ctx.Next()
}
