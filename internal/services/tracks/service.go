package tracks

import (
	"context"
	"proyek3-catalog-music/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type SpotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type Service struct {
	SpotifyOutbound SpotifyOutbound
}

func NewService(spotifyOutbound SpotifyOutbound) *Service {
	return &Service{
		SpotifyOutbound: spotifyOutbound,
	}
}
