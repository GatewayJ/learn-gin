package model

type UserPassword struct {
	Name     string `form:"username"`
	Password string `form:"password"`
}
