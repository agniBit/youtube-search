package api

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"

	"github.com/agniBit/youtube-search/pkg/gateway"
	"github.com/agniBit/youtube-search/pkg/gateway/transport"
	"github.com/agniBit/youtube-search/pkg/youtube"
	yt_repository "github.com/agniBit/youtube-search/pkg/youtube/repository/postgres"
	utils "github.com/agniBit/youtube-search/utl/common"
	"github.com/agniBit/youtube-search/utl/config"
	"github.com/agniBit/youtube-search/utl/server"
	"github.com/agniBit/youtube-search/utl/storage/postgres"
)

func Start(configPath string) error {
	// Initialize cfg
	cfg, err := config.Load(configPath)
	// check error and panic in case of error
	utils.CheckErr(err)

	// initialize gorm db
	db, err := postgres.NewGormDB(cfg.DB)
	if err != nil {
		return err
	}

	// init youtube client
	ytService, err := yt.NewService(context.Background(), option.WithAPIKey(cfg.Youtube.APIKeys[0])) // use first key on initialization
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// initialize youtube repository
	youtubeRepo := yt_repository.NewYoutubeRepository(db)

	// Initialize youtube service
	yt := youtube.New(youtubeRepo, cfg, ytService)

	// Initialize gateway service
	gw := gateway.New(yt)

	// Initialize new echo server
	e := echo.New()

	// register all APIs in echo group as version 1 APIs
	v1 := e.Group("/v1")

	// register all HTTP APIs
	transport.NewHTTP(gw, v1)

	// Start HTTP server
	server.Start(e, cfg.Server)

	return nil
}
