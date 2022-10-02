package youtube

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	youtubeType "github.com/agniBit/youtube-search/type/youtube"
	"github.com/agniBit/youtube-search/utl/config"
)

var source string = "job/youtube/youtube.go"

func New(ctx context.Context, cron *cron.Cron, ytS youtubeType.Service, cfg *config.Cron) {
	//  run Growth & Retention analytics job
	// register cron job for fetching new videos from youtube
	// _, err := cron.AddFunc(fmt.Sprintf("*/%d * * * *", cfg.Youtube.FetchNewVideosInterval), func() {
	// 	log.Printf("%s: fetching new youtube videos", source)
	// 	err := ytS.FetchNewYoutubeVideos()

	// 	if err != nil {
	// 		log.Printf("%s: error in sending growth and retention summary: %v", source, err)
	// 	}
	// })

	go func(ctx context.Context) {
		for {
			startTime := time.Now()
			log.Printf("%s: fetching new youtube videos", source)
			err := ytS.FetchNewYoutubeVideos(ctx, "cricket")

			if err != nil {
				log.Printf("%s: error in sending growth and retention summary: %v", source, err)
			}

			// total time taken to run the job
			elapsedTime := time.Since(startTime)

			// sleep for the remaining time
			sleepTime := time.Duration(cfg.Youtube.FetchNewVideosInterval)*time.Second - elapsedTime

			time.Sleep(sleepTime)
		}
	}(ctx)
}
