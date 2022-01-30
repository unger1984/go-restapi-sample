package users

import (
	domain "awcoding.com/back/src/domain/users"
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
	query := `
SELECT "User".*, avatar."id" as "avatar.id", avatar."path" "avatar.path", 
		avatar."name" "avatar.name", avatar."type" "avatar.type" 
FROM "User" 
LEFT JOIN "Upload" as avatar ON avatar."id"="User"."avatarId" 
WHERE "User"."id"=$1`
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

func (r *Repository) GetByEmailPassword(email string, password string) (*domain.User, error) {
	var user User
	query := `
SELECT "User".*, avatar."id" as "avatar.id", avatar."path" "avatar.path",
		avatar."name" "avatar.name", avatar."type" "avatar.type"
FROM "User"
LEFT JOIN "Upload" as avatar ON avatar."id"="User"."avatarId"
WHERE "User"."email"=$1 and "User"."password" ilike $2`
	rdb := r.db.Unsafe()
	if err := rdb.Get(&user, query, email, password); err != nil {
		panic(err)
		return nil, err
	}
	return user.ToEntity(), nil
}
