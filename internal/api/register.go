package api

import (
	"net/http"

	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type Regsiter struct {
	RegisterService interfaces.IRegisterService
}

func (api *Regsiter) RegisterHandler(c *gin.Context) {
	log := helpers.Logger

	req := models.User{}

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

	response, err := api.RegisterService.Register(c, req)
	if err != nil {
		log.Error("failed to register new user ", err)
		helpers.SendResponseHttp(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	helpers.SendResponseHttp(c, http.StatusOK, constants.SuccessMessage, response)
}
