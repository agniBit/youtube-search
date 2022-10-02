package gateway

import youtube_type "github.com/agniBit/youtube-search/type/youtube"

func (g GatewayService) SearchYoutubeVideosByName(name string) ([]*youtube_type.YoutubeVideo, error) {
	return g.yt.SearchYoutubeVideosByName(name)
}
