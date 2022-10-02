package youtube

import (
	"context"
	"time"

	yt "google.golang.org/api/youtube/v3"

	"github.com/agniBit/youtube-search/type/common"
	youtube_type "github.com/agniBit/youtube-search/type/youtube"
	"github.com/agniBit/youtube-search/utl/config"
	"github.com/jinzhu/copier"
)

type YoutubeService struct {
	repository Repository
	config     *config.Configuration
	ytService  *yt.Service
}

func New(repository Repository, config *config.Configuration, ytService *yt.Service) youtube_type.Service {
	return &YoutubeService{repository: repository, config: config, ytService: ytService}
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

func copyThumbnail(youtubeVideo *YoutubeVideo, ytThumbnail *yt.Thumbnail, resolutionType ThumbnailType) error {
	thumbnail := &YoutubeVideoThumbnails{}
	err := copier.Copy(thumbnail, ytThumbnail)
	if err != nil {
		return err
	}
	thumbnail.ResolutionType = resolutionType
	youtubeVideo.Thumbnails = append(youtubeVideo.Thumbnails, thumbnail)

	return nil
}

func YoutubeSearchResultToYoutubeVideo(youtubeSearchResult *yt.SearchResult) (*YoutubeVideo, error) {
	youtubeVideo := &YoutubeVideo{}
	err := copier.Copy(youtubeVideo, youtubeSearchResult)
	if err != nil {
		return nil, err
	}

	if youtubeSearchResult.Snippet != nil {
		err = copier.Copy(youtubeVideo, youtubeSearchResult.Snippet)
		if err != nil {
			return nil, err
		}

		youtubeVideo.Thumbnails = nil
		err := copyThumbnail(youtubeVideo, youtubeSearchResult.Snippet.Thumbnails.Default, ThumbnailTypeDefault)
		if err != nil {
			return nil, err
		}
		err = copyThumbnail(youtubeVideo, youtubeSearchResult.Snippet.Thumbnails.High, ThumbnailTypeHigh)
		if err != nil {
			return nil, err
		}
		err = copyThumbnail(youtubeVideo, youtubeSearchResult.Snippet.Thumbnails.Medium, ThumbnailTypeMedium)
		if err != nil {
			return nil, err
		}

		err = copier.Copy(youtubeVideo, youtubeSearchResult.Id)
		if err != nil {
			return nil, err
		}

		youtubeVideo.PublishedAt, err = time.Parse(time.RFC3339, youtubeSearchResult.Snippet.PublishedAt)
		if err != nil {
			return nil, err
		}
		youtubeVideo.ChannelId = youtubeSearchResult.Snippet.ChannelId
	}

	return youtubeVideo, nil
}

func (yt *YoutubeService) FetchNewYoutubeVideos(ctx context.Context, search string) error {
	// Make the API call to YouTube.
	call := yt.ytService.Search.List([]string{"id", "snippet"}).Q(search).MaxResults(100)
	response, err := call.Do()
	if err != nil {
		return err
	}

	youtubeVideos := []*YoutubeVideo{}

	for _, resp := range response.Items {
		video, err := YoutubeSearchResultToYoutubeVideo(resp)
		if err != nil {
			return err
		}
		youtubeVideos = append(youtubeVideos, video)
	}

	err = yt.repository.SaveYoutubeVideos(ctx, youtubeVideos)
	return err
}
