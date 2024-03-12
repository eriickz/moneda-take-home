package auth

import "github.com/labstack/echo/v4"

func RegisterAuthRoutes(apiGroup *echo.Group) {
  authGroup := apiGroup.Group("/auth")

  authGroup.POST("/sign-in", SignIn)
}
