package youtube

import (
	"github.com/labstack/echo/v4"

	"github.com/agniBit/youtube-search/pkg/gateway"
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
	name := c.QueryParam("name")

	videos, err := h.svc.SearchYoutubeVideosByName(name)
	if err != nil {
		return err
	}

	return c.JSON(200, videos)
}
