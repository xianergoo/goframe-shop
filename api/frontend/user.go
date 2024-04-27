package frontend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type RegisterUserReq struct {
	g.Meta       `path:"/register" method:"post" tags:"front" summary:"register user"`
	Name         string `json:"name"         description:"用户名" v:"required#Name is required"`
	Avatar       string `json:"avatar"       description:"头像"`
	Password     string `json:"password"     description:"" v:"password"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterUserRes struct {
	Id uint `json:"id"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	//Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}

// for jwt
type LoginRes struct {
	//Info interface{} `json:"info"`
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
