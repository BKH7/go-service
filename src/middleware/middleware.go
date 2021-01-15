package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Middleware struct
type Middleware struct{}

// CORS will handle the CORS middleware
func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		return next(c)
	}
}

//RemoveTrailingSlash from the request URI.
func (m *Middleware) RemoveTrailingSlash() (midFunc echo.MiddlewareFunc) {
	return middleware.RemoveTrailingSlash()
}

// Initial middleware when called
func Initial() *Middleware {
	return &Middleware{}
}
