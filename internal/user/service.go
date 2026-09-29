package user

import (
	"fmt"
	"gotickets/internal/user/dto"
)

var ErrorInvalidCredentials = fmt.Errorf("Invalid email or password!")

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req dto.CreateUserRequestDto) (*dto.UserResponseDto, error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		// Password: req.Password,
	}
	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseDto{
		ID: 	  user.ID,
		Name:     user.Name,
		Email:    user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return response, nil
}

func (s *service) LoginUser(req dto.LoginUserRequestDto) (*dto.UserResponseDto, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrorInvalidCredentials
	}

	err = user.checkPassword(req.Password)
	if err != nil {
		return nil, ErrorInvalidCredentials
	}

	response := dto.UserResponseDto{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}
