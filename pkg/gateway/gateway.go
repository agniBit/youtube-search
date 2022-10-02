package gateway

import (
	youtubeType "github.com/agniBit/youtube-search/type/youtube"
)

type GatewayService struct {
	yt youtubeType.Service
}

// New creates new gateway service
func New(
	yt youtubeType.Service,
) Service {
	return GatewayService{
		yt: yt,
	}
}
