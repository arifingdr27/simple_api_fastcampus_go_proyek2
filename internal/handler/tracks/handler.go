package tracks

import (
	"context"
	spotifymodel "proyek3-catalog-music/internal/models/spotify"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageLimit, pageIndex int) (*spotifymodel.SearchResponse, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, s service) *Handler {
	return &Handler{
		Engine:  api,
		service: s,
	}
}

func (h *Handler) RegisterRoutes() {
	tracks := h.Group("/tracks")
	tracks.GET("/search", h.Search)
}
