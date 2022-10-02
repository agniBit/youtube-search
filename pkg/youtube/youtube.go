package youtube

import (
	"fmt"

	youtube_type "github.com/agniBit/youtube-search/type/youtube"
	"github.com/agniBit/youtube-search/utl/config"
)

type YoutubeService struct {
	config *config.Configuration
}

func New(config *config.Configuration) youtube_type.Service {
	return &YoutubeService{config: config}
}

func (yt *YoutubeService) SearchYoutubeVideosByName(name string) ([]*youtube_type.YoutubeVideo, error) {
	return nil, fmt.Errorf("not implemented")
}
