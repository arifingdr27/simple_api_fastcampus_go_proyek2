package memberships

import (
	"net/http"
	"proyek3-catalog-music/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var login memberships.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	access_token, err := h.service.Login(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, memberships.LoginResponse{
		AccessToken: access_token,
	})
}
