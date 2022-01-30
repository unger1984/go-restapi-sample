package entities

type User struct {
	Id       int     `json:"id"`
	Email    string  `json:"email"`
	AvatarId *int    `json:"avatarId"`
	Avatar   *Upload `json:"avatar"`
	Password string  `json:"-"`
}
