package interfaces

import "github.com/gin-gonic/gin"

type IRegisterHandler interface {
	RegisterHandler(c *gin.Context)
}
