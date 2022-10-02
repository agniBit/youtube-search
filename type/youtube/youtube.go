package youtubeType

import (
	"context"
	"time"

	"github.com/agniBit/youtube-search/type/common"
)

type (
	Service interface {
		FetchNewYoutubeVideos(crx context.Context, search string) error
		SearchYoutubeVideosByName(ctx context.Context, search *SearchFilter, offsetLimit *common.OffsetLimit) ([]*YoutubeVideo, error)
	}

	YoutubeVideo struct {
		ID   string `json:"id"`
		Etag string `json:"etag,omitempty"`
		Kind string `json:"kind,omitempty"`
		*YoutubeVideoID
		*YoutubeVideoSnippet
	}

	YoutubeVideoID struct {
		Kind    string `json:"kind,omitempty"`
		VideoId string `json:"videoId,omitempty"`
	}

	YoutubeVideoSnippet struct {
		ChannelId            string                    `json:"channelId"`
		Title                string                    `json:"title"`
		Description          string                    `json:"description"`
		ChannelTitle         string                    `json:"channelTitle"`
		LiveBroadcastContent string                    `json:"liveBroadcastContent"`
		Thumbnails           []*YoutubeVideoThumbnails `json:"thumbnails"`
		PublishedAt          time.Time                 `json:"publishedAt"`
	}

	YoutubeVideoThumbnails struct {
		ID             string        `json:"id"`
		ResolutionType ThumbnailType `json:"resolutionType"`
		Height         int           `json:"height"`
		Url            string        `json:"url"`
		Width          int           `json:"width"`
		YoutubeVideoID string        `json:"youtubeVideoId"`
		YoutubeVideo   *YoutubeVideo `json:"youtubeVideo"`
	}

	SearchFilter struct {
		Title       string `query:"title"`
		Description string `query:"description"`
		Search      string `query:"search"`
	}

	ThumbnailType string
)
