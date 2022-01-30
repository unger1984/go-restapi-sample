package users

type Repository interface {
	GetById(id int) (*User, error)
	GetByEmailPassword(email string, password string) (*User, error)
}
