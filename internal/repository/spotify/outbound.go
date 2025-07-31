package spotify

import (
	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/pkg/httpclient"
	"time"
)

type Outbound struct {
	cfg         configs.Config
	client      httpclient.HTTPClient
	accessToken string
	tokenType   string
	expiredTime time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, client httpclient.HTTPClient) *Outbound {
	return &Outbound{
		cfg:    *cfg,
		client: client,
	}
}
