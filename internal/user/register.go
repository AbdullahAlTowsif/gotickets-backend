package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
}
