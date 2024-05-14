package user

import (
	"github.com/onainadapdap1/golang_kantin/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	GetUserByID(userID int) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(userID int) (models.User, error) {
	tx := r.db.Begin()
	var user models.User

	if err := tx.Debug().Where("id = ?", userID).First(&user).Error; err != nil {
		return models.User{}, nil
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	tx := r.db.Begin()
	var user models.User
	if err := tx.Debug().Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
