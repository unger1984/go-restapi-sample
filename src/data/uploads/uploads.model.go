package uploads

import (
	domain "awcoding.com/back/src/domain/uploads"
)

type Upload struct {
	Id   *int    `db:"id"`
	Path *string `db:"path"`
	Name *string `db:"name"`
	Type *string `db:"type"`
}

func (u *Upload) ToEntity() *domain.Upload {
	return &domain.Upload{
		Id:   u.Id,
		Path: u.Path,
		Name: u.Name,
		Type: u.Type,
	}
}
