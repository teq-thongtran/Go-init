package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4/middleware"

	"myapp/http/card"
	"myapp/http/user"
	"myapp/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOriginFunc: func(origin string) (bool, error) {
			return regexp.MatchString(
				`^https:\/\/(|[a-zA-Z0-9]+[a-zA-Z0-9-._]*[a-zA-Z0-9]+\.)teqnological.asia$`,
				origin,
			)
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	}))

	// APIs
	api := e.Group("/api")
	userApi := api.Group("/users/:user_id")
	user.Init(api.Group("/users"), useCase)
	card.Init(userApi.Group("/cards"), useCase)
	return e
}
