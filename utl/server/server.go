package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/agniBit/youtube-search/utl/config"
)

func Start(e *echo.Echo, cfg *config.Server) {
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
	}
	e.Debug = cfg.Debug

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	// wait for some time to finish the old requests
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
