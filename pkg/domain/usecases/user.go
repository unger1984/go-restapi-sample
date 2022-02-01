package usecases

import (
	"awcoding.com/back/pkg/domain/entities"
)

//go:generate mockgen --source=user.go --destination=mocks/user_mock.go

// UserRepository is repository template
type UserRepository interface {
	GetById(id int) (*entities.User, error)
	GetByEmailPassword(email string, password string) (*entities.User, error)
}

// UserCases is usecase interfece duplicate repository
type UserCases interface {
	GetById(id int) (*entities.User, error)
	GetByEmailPassword(email string, password string) (*entities.User, error)
}

type userCases struct {
	repository UserCases
}

func NewUserCases(repository UserRepository) *userCases {
	return &userCases{repository: repository}
}

func (s *userCases) GetById(id int) (*entities.User, error) {
	return s.repository.GetById(id)
}

func (s *userCases) GetByEmailPassword(email string, password string) (*entities.User, error) {
	return s.repository.GetByEmailPassword(email, password)
}
