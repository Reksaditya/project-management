package services

import (
	"errors"

	"github.com/Reksaditya/project-management/models"
	"github.com/Reksaditya/project-management/repositories"
	"github.com/Reksaditya/project-management/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserSevice(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)

	if existingUser.UserID != 0 {
		return errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()
	return s.repo.Create(user)
}