package users

type UserService interface {
	GetById(id int) (*User, error)
	GetByEmailPassword(email string, password string) (*User, error)
}

type Service struct {
	repository UserService
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetById(id int) (*User, error) {
	return s.repository.GetById(id)
}

func (s *Service) GetByEmailPassword(email string, password string) (*User, error) {
	return s.repository.GetByEmailPassword(email, password)
}
