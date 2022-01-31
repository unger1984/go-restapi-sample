package models

import (
	"awcoding.com/back/pkg/domain/entities"
	"time"
)

type User struct {
	Id        int     `gorm:"primaryKey;column:id"`
	Email     string  `gorm:"column:email"`
	AvatarId  *int    `gorm:"column:avatarId" `
	Password  string  `gorm:"column:password"`
	Avatar    *Upload `gorm:"foreignKey:AvatarId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "User"
}

func (u User) ToEntity() *entities.User {
	avatar := u.Avatar
	var avatarEntity *entities.Upload
	if avatar != nil {
		avatarEntity = avatar.ToEntity()
	}
	return &entities.User{Id: u.Id, Email: u.Email, Password: u.Password, AvatarId: u.AvatarId, Avatar: avatarEntity}
}
