package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	pageSize := c.Query("pageSize")
	pageIndex := c.Query("pageIndex")

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil {
		pageIndexInt = 1
	}

	response, err := h.service.Search(ctx, query, pageSizeInt, pageIndexInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
