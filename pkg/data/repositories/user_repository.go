package repositories

import (
	"awcoding.com/back/pkg/data/models"
	"awcoding.com/back/pkg/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(id int) (*entities.User, error) {
	var user models.User
	if result := r.db.Joins("Avatar").First(&user, id); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected <= 0 {
		return nil, errors.New("user not found")
	}

	return user.ToEntity(), nil
}

func (r *UserRepository) GetByEmailPassword(email string, password string) (*entities.User, error) {
	var user models.User
	if result := r.db.Joins("Avatar").Find(&user, "email=? and password iLike ?", email, password); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected <= 0 {
		return nil, errors.New("login and password incorrect")
	}

	return user.ToEntity(), nil
}
