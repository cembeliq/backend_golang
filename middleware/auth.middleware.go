package middleware

import "github.com/labstack/echo/v4/middleware"

// IsAuthenticated is middleware
var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secretadmirer"),
})
