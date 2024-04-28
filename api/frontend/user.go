package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterUserReq struct {
	g.Meta       `path:"/register" method:"post" tags:"front" summary:"user register"`
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
	g.Meta   `path:"/login" method:"post" tags:front"" summary:"user login"`
	Name     string `json:"passport" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"password"   dc:"密码(明文)"`
}

// for gtoken
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sex      uint8  `json:"sex"`
	Sign     string `json:"sign"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"front" summary:"user info"`
}

// for gtoken
type UserInfoRes struct {
	UserInfoBase
}

// 可以复用的，一定要抽取出来
type UserInfoBase struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

type UpdatePasswordReq struct {
	g.Meta       `path:"/update/password" method:"post" tags:"front" summary:"modify password"`
	Password     string `json:"password" v:"password" dc:"password"`
	UserSalt     string `json:"userSalt,omitempty" dc:"加密盐"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type UpdatePasswordRes struct {
	Id uint `json:"id"`
}
