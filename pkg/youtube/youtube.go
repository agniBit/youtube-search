package youtube

import (
	"context"

	"github.com/agniBit/youtube-search/type/common"
	youtube_type "github.com/agniBit/youtube-search/type/youtube"
	"github.com/agniBit/youtube-search/utl/config"
	"github.com/jinzhu/copier"
)

type YoutubeService struct {
	repository Repository
	config     *config.Configuration
}

func New(repository Repository, config *config.Configuration) youtube_type.Service {
	return &YoutubeService{repository: repository, config: config}
}

func (yt *YoutubeService) SearchYoutubeVideosByName(ctx context.Context, search *youtube_type.SearchFilter, offsetLimit *common.OffsetLimit) ([]*youtube_type.YoutubeVideo, error) {
	videosT := []*youtube_type.YoutubeVideo{}

	videos, err := yt.repository.FindVideosByVideoName(search, offsetLimit)
	if err != nil {
		return nil, err
	}

	// copy data in DTO for http transport
	err = copier.Copy(&videosT, &videos)
	return videosT, err
}

func (yt *YoutubeService) FetchNewYoutubeVideos(ctx context.Context) error {
	return nil
}
