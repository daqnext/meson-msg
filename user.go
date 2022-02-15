package meson_msg

type ValidateToken struct {
	Token string
}

type LoginMsg struct {
	UserName  string
	Email     string
	Password  string `binding:"required"`
	Captcha   string `binding:"required"`
	CaptchaId string `binding:"required"`
	Type      string
}

type RegisterMsg struct {
	UserName  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Vcode     string `binding:"required"`
	UserType  []string
	Captcha   string `binding:"required""`
	CaptchaId string `binding:"required"`
}
