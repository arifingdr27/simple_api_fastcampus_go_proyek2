package api

import (
	"net/http"

	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api LoginHandler) LoginHandlerHttp(c *gin.Context) {
	log := helpers.Logger
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHttp(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
		return
	}

	resp, err := api.LoginService.Login(c, req)
	if err != nil {
		log.Error("failed to login: ", err)
		helpers.SendResponseHttp(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, resp)
}
