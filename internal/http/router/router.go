package router

import (
	"ablufus/internal/database"
	"ablufus/internal/http/handler"
	"ablufus/internal/http/health"
	"ablufus/internal/repository"
	"ablufus/internal/service"

	"github.com/labstack/echo/v4"
)

func Handlers(db database.Database) *echo.Echo {
	e := echo.New()

	r := repository.New(db)
	s := service.New(r)
	h := handler.New(s)

	e.GET("/health", health.Health)

	v1 := e.Group("/v1/user")
	v1.POST("", h.Post)
	v1.GET("", h.List)
	v1.PATCH("/:user_id", h.Update)

	return e
}
