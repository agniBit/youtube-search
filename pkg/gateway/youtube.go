package gateway

import (
	"context"

	"github.com/agniBit/youtube-search/type/common"
	youtube_type "github.com/agniBit/youtube-search/type/youtube"
)

func (g GatewayService) SearchYoutubeVideosByName(ctx context.Context, search *youtube_type.SearchFilter, offsetLimit *common.OffsetLimit) ([]*youtube_type.YoutubeVideo, error) {
	return g.yt.SearchYoutubeVideosByName(ctx, search, offsetLimit)
}
