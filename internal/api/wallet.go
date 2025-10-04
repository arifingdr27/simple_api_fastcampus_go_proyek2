package api

import (
	"net/http"

	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type WalletAPI struct {
	WalletService interfaces.IWalletService
}

func (api *WalletAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Wallet
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if req.UserID == 0 {
		log.Error("user_id is empty")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err := api.WalletService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, req)
}
