package user

import (
	"errors"

	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/repository/user"
	"github.com/onainadapdap1/golang_kantin/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(input api.LoginInput) (models.User, error)
	GetUserByID(userID int) (models.User, error)
}

type userService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(userID int) (models.User, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}
	return user, nil
}

func (s *userService) Login(input api.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	if user.ID == 0 {
		return user, errors.New("tidak ada user dengan email tersebut")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
