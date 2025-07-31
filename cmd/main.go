package main

import (
	"log"
	"net/http"

	"proyek3-catalog-music/internal/configs"
	membershipHandler "proyek3-catalog-music/internal/handler/memberships"
	trackHandler "proyek3-catalog-music/internal/handler/tracks"
	"proyek3-catalog-music/internal/models/memberships"
	membershipRepo "proyek3-catalog-music/internal/repository/memberships"
	"proyek3-catalog-music/internal/repository/spotify"
	membershipService "proyek3-catalog-music/internal/services/memberships"
	"proyek3-catalog-music/internal/services/tracks"
	"proyek3-catalog-music/pkg/httpclient"
	"proyek3-catalog-music/pkg/internalsql"

	"github.com/gin-gonic/gin"
)

func main() {
	var cfg configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/",
		}),
		configs.WithConfigFile("config.yaml"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	cfg = *configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&memberships.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	r := gin.Default()

	httpclient := httpclient.NewClient(&http.Client{})
	spotifyOutbound := spotify.NewSpotifyOutbound(&cfg, httpclient)

	memberRepository := membershipRepo.NewRepository(db)
	membershipService := membershipService.NewService(&cfg, memberRepository)
	trackService := tracks.NewService(spotifyOutbound)

	membershipHandler := membershipHandler.NewHandler(r, membershipService)
	membershipHandler.RegisterRoutes()
	trackHandler := trackHandler.NewHandler(r, trackService)
	trackHandler.RegisterRoutes()
	r.Run(cfg.Service.Port) // Start the server on the configured address
}
