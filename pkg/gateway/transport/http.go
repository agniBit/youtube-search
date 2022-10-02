package transport

import (
	"github.com/labstack/echo/v4"

	"github.com/agniBit/youtube-search/pkg/gateway"
	"github.com/agniBit/youtube-search/pkg/gateway/transport/youtube"
)

func NewHTTP(svc gateway.Service, e *echo.Group) {
	// register youtube APIs in echo group
	youtube.NewHTTP(svc, e)
}
