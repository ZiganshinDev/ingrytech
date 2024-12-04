package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Handler interface {
	URLs() map[string]map[string]echo.HandlerFunc
}

type Server struct {
	h *http.Server
}

func New(port string, handler Handler, logger *zap.SugaredLogger) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			fmt.Sprintf("http://localhost:%s", port),
		},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowCredentials: true,
	}))

	e.Use(recoveryMiddleware(logger))
	e.HTTPErrorHandler = (httpErrorHandler(logger))

	internalApi := e.Group("/api/v1")

	addURLs(internalApi, handler.URLs())

	h := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: e,
	}

	return &Server{h: h}
}

func addURLs(group *echo.Group, handlers map[string]map[string]echo.HandlerFunc) {
	for path := range handlers {
		for method, handler := range handlers[path] {
			group.Add(method, path, handler)
		}
	}
}

func (s *Server) Start() error {
	if err := s.h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.h.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
