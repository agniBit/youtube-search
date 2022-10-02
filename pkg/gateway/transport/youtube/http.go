package youtube

import (
	"github.com/labstack/echo/v4"

	"github.com/agniBit/youtube-search/pkg/gateway"
	"github.com/agniBit/youtube-search/type/common"
	youtubeType "github.com/agniBit/youtube-search/type/youtube"
)

type HTTP struct {
	svc gateway.Service
}

func NewHTTP(svc gateway.Service, e *echo.Group) {
	h := HTTP{svc: svc}

	youtubeGrp := e.Group("/youtube")

	// APIs
	youtubeGrp.GET("/search", h.searchYoutubeVideos)
}

func (h *HTTP) searchYoutubeVideos(c echo.Context) error {
	search := &youtubeType.SearchFilter{}
	if err := c.Bind(search); err != nil {
		return err
	}

	offsetLimit := &common.OffsetLimit{}
	if err := c.Bind(offsetLimit); err != nil {
		return err
	}

	videos, err := h.svc.SearchYoutubeVideosByName(c.Request().Context(), search, offsetLimit)
	if err != nil {
		return err
	}

	return c.JSON(200, videos)
}
