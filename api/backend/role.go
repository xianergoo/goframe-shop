package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" tags:"Role" desc:"add role"`
	Name   string `json:"name" v:"required"#require name" dc:"role name"`
	Desc   string `json:"desc" desc:"role description"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update/{Id}" method:"post" tags:"Role" summary:"修改role接口"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的role" dc:"id"`
	Name   string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Desc   string `json:"desc"    dc:"角色ids"`
}
type RoleUpdateRes struct {
	Id uint `json:"id"`
}

type RoleInfoReq struct {
	g.Meta `path:"/backend/role/info/{Id}" method:"get" tags:"Role" summary:"get info"`
	Id     uint `json:"id"      v:"min:1#请选择需要查看的role" dc:"id"`
}
type RoleInfoRes struct {
	Id   uint   `json:"id"     dc:"id"`
	Name string `json:"name" dc:"用户名"`
	Desc string `json:"desc"   dc:"角色ids"`
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"Role" summary:"role列表接口"`
	CommonPaginationReq
}

type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"Role" summary:"删除Role接口"`
	Id     uint `v:"min:1#请选择需要删除的Role" dc:"Roleid"`
}

type RoleDeleteRes struct{}

type RoleAddPermssionReq struct {
	g.Meta       `path:"/backend/role/add/permission" method:"post" tags:"Role" summary:"add permission"`
	RoleId       uint `json:"role_id" desc:"role id"`
	PermissionId uint `json:"permission_id" desc:"role permission id"`
}

type RoleAddPermssionRes struct {
	Id uint `json:"id" desc:"role id"`
}

type RoleDelPermssionReq struct {
	g.Meta       `path:"/backend/role/delete/permission" method:"post" tags:"Role" summary:"delete permission"`
	RoleId       uint `json:"role_id" desc:"role id"`
	PermissionId uint `json:"permission_id" desc:"role permission id"`
}

type RoleDelPermssionRes struct {
}
