package users

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
