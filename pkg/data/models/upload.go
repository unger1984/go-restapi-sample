package models

import (
	"awcoding.com/back/pkg/domain/entities"
	"time"
)

type Upload struct {
	Id        int    `gorm:"primaryKey;column:id"`
	Path      string `db:"column:path"`
	Name      string `db:"column:name"`
	Type      string `db:"column:type"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Upload) TableName() string {
	return "Upload"
}

func (u *Upload) ToEntity() *entities.Upload {
	return &entities.Upload{
		Id:   u.Id,
		Path: u.Path,
		Name: u.Name,
		Type: u.Type,
	}
}
