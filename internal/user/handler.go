package user

import (
	"errors"
	"gotickets/internal/httpResponse"
	"gotickets/internal/user/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateUserRequestDto
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation Failed!",
			Details: err.Error(),
		})
	}

	res, err := h.service.CreateUser(req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExists) {
			return c.JSON(http.StatusConflict, httpResponse.Error {
				Code: http.StatusConflict,
				Message: "Failed to create user",
				Details: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginUserRequestDto
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation Failed!",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(req)
	if err != nil  {
		if errors.Is(err, ErrorInvalidCredentials) {
			return c.JSON(http.StatusUnauthorized, httpResponse.Error {
				Code: http.StatusUnauthorized,
				Message: "Invalid email or password",
				Details: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Code: http.StatusInternalServerError,
			Message: "Failed to login user",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}


func (h *handler) GetMe(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpResponse.Error{
			Code: http.StatusUnauthorized,
			Message: "Cannot get user information",
			Details: "Missing user_id in context",
		})
	}

	email, _ := c.Get("user_email").(string)
	name, _ := c.Get("user_name").(string)

	return c.JSON(http.StatusOK, dto.UserResponseDto{
		ID: userID,
		Name: name,
		Email: email,
	})
}
