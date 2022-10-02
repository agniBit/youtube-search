package youtubeType

import (
	"context"

	"github.com/agniBit/youtube-search/type/common"
)

type (
	Service interface {
		SearchYoutubeVideosByName(ctx context.Context, search *SearchFilter, offsetLimit *common.OffsetLimit) ([]*YoutubeVideo, error)
	}

	YoutubeVideo struct {
		ID           string `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ThumbnailURL string `json:"thumbnail_url"`
		URL          string `json:"url"`
		PublishedAt  string `json:"published_at"`
	}

	SearchFilter struct {
		Title       string `query:"title"`
		Description string `query:"description"`
		Search      string `query:"search"`
	}
)
