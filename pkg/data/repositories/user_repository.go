package repositories

import (
	"awcoding.com/back/pkg/data/models"
	"awcoding.com/back/pkg/domain/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(id int) (*entities.User, error) {
	var user models.User
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

func (r *UserRepository) GetByEmailPassword(email string, password string) (*entities.User, error) {
	var user models.User
	query := `
SELECT "User".*, avatar."id" as "avatar.id", avatar."path" "avatar.path",
		avatar."name" "avatar.name", avatar."type" "avatar.type"
FROM "User"
LEFT JOIN "Upload" as avatar ON avatar."id"="User"."avatarId"
WHERE "User"."email"=$1 and "User"."password" ilike $2`
	rdb := r.db.Unsafe()
	if err := rdb.Get(&user, query, email, password); err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}
