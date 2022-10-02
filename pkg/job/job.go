package job

import (
	"context"
	"log"

	"github.com/robfig/cron/v3"
	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"

	"github.com/agniBit/youtube-search/pkg/job/youtube"
	youtubeS "github.com/agniBit/youtube-search/pkg/youtube"
	yt_repository "github.com/agniBit/youtube-search/pkg/youtube/repository/postgres"
	utils "github.com/agniBit/youtube-search/utl/common"
	"github.com/agniBit/youtube-search/utl/config"
	"github.com/agniBit/youtube-search/utl/storage/postgres"
)

func Start(configPath string) error {
	c := cron.New()

	cfg, err := config.Load(configPath)
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

	// Initialize youtube service
	ytS := youtubeS.New(yt_repository.NewYoutubeRepository(db), cfg, ytService)

	// register cron job for fetching new videos from youtube
	youtube.New(context.Background(), c, ytS, cfg.Cron)

	c.Start()
	defer c.Stop()

	// block indefinitely
	select {}
}
