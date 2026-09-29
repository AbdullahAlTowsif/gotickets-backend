package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	e.POST("/users", userHandler.CreateUser)
}
