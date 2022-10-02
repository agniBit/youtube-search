package job

import (
	"context"

	"github.com/robfig/cron/v3"

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

	// Initialize youtube service
	ytS := youtubeS.New(yt_repository.NewYoutubeRepository(db), cfg)

	// register cron job for fetching new videos from youtube
	youtube.New(context.Background(), c, ytS, cfg.Cron)

	c.Start()
	defer c.Stop()

	// block indefinitely
	select {}
}
