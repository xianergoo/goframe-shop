package backend

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	//Captcha  string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
}
type LoginDoRes struct {
	Info interface{}
	//???
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
}
