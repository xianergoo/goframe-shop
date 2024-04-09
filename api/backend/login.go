package backend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	//Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type LoginDoRes struct {
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post"`
}

type LogoutRes struct {
}
