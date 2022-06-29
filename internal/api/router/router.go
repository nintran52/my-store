package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nintran52/my-store/internal/api"
	"github.com/nintran52/my-store/internal/api/handlers"
	"github.com/nintran52/my-store/internal/api/middleware"
)

func Init(s *api.Server) {
	s.Echo = echo.New()

	s.Router = &api.Router{
		Routes:    nil,
		APIV1Auth: s.Echo.Group("/api/v1/auth", middleware.AuthWithConfig()),
	}

	handlers.AttachAllRoutes(s)
}
