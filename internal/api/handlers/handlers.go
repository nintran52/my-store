package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/nintran52/my-store/internal/api"
)

func AttachAllRoutes(s *api.Server) {
	s.Router.Routes = []*echo.Route{}
}
