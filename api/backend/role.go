package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" tags:"role" desc:"role add"`
	Name   string `json:"name" v:"required"#require name" dc:"role name"`
	Desc   string `json:"desc" desc:"role description"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update/{Id}" method:"post" tags:"role" summary:"role modify"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的role" dc:"id"`
	Name   string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Desc   string `json:"desc"    dc:"角色ids"`
}
type RoleUpdateRes struct {
	Id uint `json:"id"`
}

type RoleInfoReq struct {
	g.Meta `path:"/role/info/{Id}" method:"get" tags:"role" summary:"role info"`
	Id     uint `json:"id"      v:"min:1#请选择需要查看的role" dc:"id"`
}
type RoleInfoRes struct {
	Id   uint   `json:"id"     dc:"id"`
	Name string `json:"name" dc:"用户名"`
	Desc string `json:"desc"   dc:"角色ids"`
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"role" summary:"role list"`
	CommonPaginationReq
}

type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"role" summary:"role delete"`
	Id     uint `v:"min:1#请选择需要删除的Role" dc:"Roleid"`
}

type RoleDeleteRes struct{}

type RoleAddPermssionReq struct {
	g.Meta       `path:"/role/add/permission" method:"post" tags:"role permission" summary:"permission add "`
	RoleId       uint `json:"role_id" desc:"role id"`
	PermissionId uint `json:"permission_id" desc:"role permission id"`
}

type RoleAddPermssionRes struct {
	Id uint `json:"id" desc:"role id"`
}

type RoleDelPermssionReq struct {
	g.Meta       `path:"/role/delete/permission" method:"post" tags:"role permission" summary:"permission delete"`
	RoleId       uint `json:"role_id" desc:"role id"`
	PermissionId uint `json:"permission_id" desc:"role permission id"`
}

type RoleDelPermssionRes struct {
}
