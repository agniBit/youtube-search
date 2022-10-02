package gateway

import (
	youtube_type "github.com/agniBit/youtube-search/type/youtube"
)

type GatewayService struct {
	yt youtube_type.Service
}

// New creates new gateway service
func New(
	yt youtube_type.Service,
) Service {
	return GatewayService{
		yt: yt,
	}
}
