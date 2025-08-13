package api

import (
	"net/http"

	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) LogoutHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if err := api.LogoutService.Logout(c, token); err != nil {
		helpers.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, nil)
}
