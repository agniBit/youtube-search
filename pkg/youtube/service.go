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
		ID   string `gorm:"primaryKey;default:('yvd_' || generate_uid(12));not null"`
		Etag string `gorm:"not null"`
		Kind string
		*YoutubeVideoID
		*YoutubeVideoSnippet
		Thumbnails []*YoutubeVideoThumbnails
	}

	YoutubeVideoID struct {
		Kind    string `gorm:"not null"`
		VideoId string `gorm:"not null"`
	}

	YoutubeVideoSnippet struct {
		ChannelId            string
		ChannelTitle         string
		Description          string `gorm:"not null"`
		LiveBroadcastContent string
		PublishedAt          time.Time `gorm:"index"`
		Thumbnails           []*YoutubeVideoThumbnails
		Title                string `gorm:"not null"`
	}

	YoutubeVideoThumbnails struct {
		ID             string        `gorm:"default:('yvt_' || generate_uid(12))"`
		ResolutionType ThumbnailType `gorm:"not null"`
		Height         int
		Url            string
		Width          int
		YoutubeVideoID string `gorm:"not null"`
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
