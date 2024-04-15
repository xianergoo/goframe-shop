package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" tags:"role" desc:"add role"`
	Name   string `json:"name" v:"required"#require name" dc:"role name"`
	Desc   string `json:"desc" desc:"role description"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}
