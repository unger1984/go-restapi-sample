package users

import (
	domain "awcoding.com/back/domain/users"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetById(id int) (*domain.User, error) {
	var user User
	query := fmt.Sprintf("SELECT u.id, u.email, u.password FROM %s as u WHERE u.id=$1", "users")
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return user.toEntity(), nil
}

func (r *Repository) GetByEmailPassword(email string, password string) (*domain.User, error) {
	var user User
	query := fmt.Sprintf("SELECT u.id, u.email, u.password FROM %s as u WHERE u.email=$1 and u.password ilike $2", "users")
	if err := r.db.Get(&user, query, email, password); err != nil {
		return nil, err
	}
	return user.toEntity(), nil
}
