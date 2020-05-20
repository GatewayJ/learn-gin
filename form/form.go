package form

type LoginForm struct {
	User string `form:"name" json:"name" binding:"required"`
	//这里有做简单验证，表示参数是必须的
	Password string `form:"password" json:"password" binding:"required"`
}
