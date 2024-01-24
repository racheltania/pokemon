package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EnableCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Add your frontend origin(s)
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderContentType},
	})
}
