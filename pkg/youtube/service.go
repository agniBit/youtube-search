package youtube

import (
	"context"
	"time"

	"github.com/agniBit/youtube-search/type/common"
	youtubeType "github.com/agniBit/youtube-search/type/youtube"
)

type (
	Repository interface {
		FindVideosByVideoName(search *youtubeType.SearchFilter, offsetLimit *common.OffsetLimit) ([]*YoutubeVideo, error)
		SaveYoutubeVideos(ctx context.Context, videos []*YoutubeVideo) error
	}

	YoutubeVideoSearchResponse struct {
		Id      *YoutubeVideoID `gorm:"_"`
		Kind    string
		Etag    string
		Snippet *YoutubeVideoSnippet `gorm:"_"`
	}

	YoutubeVideo struct {
		ID   string `gorm:"default:('yvd_' || generate_uid(12))"`
		Etag string `gorm:"not null"`
		Kind string
		*YoutubeVideoID
		*YoutubeVideoSnippet
	}

	YoutubeVideoID struct {
		Kind    string `gorm:"not null"`
		VideoId string `gorm:"not null"`
	}

	YoutubeVideoSnippet struct {
		ChannelId            string `json:"ChannelId"`
		ChannelTitle         string
		Description          string
		LiveBroadcastContent string
		PublishedAt          time.Time
		Thumbnails           []*YoutubeVideoThumbnails
		Title                string
	}

	Thumbnail struct {
		Height int
		Url    string
		Width  int
	}

	YoutubeVideoThumbnails struct {
		ID             string `gorm:"default:('yvt_' || generate_uid(12))"`
		ResolutionType ThumbnailType
		Height         int
		Url            string
		Width          int
		YoutubeVideoID string
		YoutubeVideo   *YoutubeVideo
	}

	ThumbnailType string
)

const (
	ThumbnailTypeDefault ThumbnailType = "default"
	ThumbnailTypeHigh    ThumbnailType = "high"
	ThumbnailTypeMedium  ThumbnailType = "medium"
)

func (YoutubeVideo) TableName() string {
	return "youtube.youtube_video"
}

func (YoutubeVideoThumbnails) TableName() string {
	return "youtube.youtube_video_thumbnails"
}
