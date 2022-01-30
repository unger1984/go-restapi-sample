package models

import (
	"awcoding.com/back/pkg/domain/entities"
)

type Upload struct {
	Id   *int    `db:"id"`
	Path *string `db:"path"`
	Name *string `db:"name"`
	Type *string `db:"type"`
}

func (u *Upload) ToEntity() *entities.Upload {
	return &entities.Upload{
		Id:   u.Id,
		Path: u.Path,
		Name: u.Name,
		Type: u.Type,
	}
}
