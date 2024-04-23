package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionReq struct {
	g.Meta `path:"/permission/add" method:"post" tags:"permission" summary:"permission add"`
	Name   string `json:"name" v:"required"#requirename" dc:"permission name"`
	Path   string `json:"path" path:"permission pathription"`
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/permission/update/{Id}" method:"post" tags:"permission" summary:"permission modify"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的permission" dc:"id"`
	PermissionCreateUpdateBase
}
type PermissionUpdateRes struct {
	Id uint `json:"id"`
}

type PermissionCreateUpdateBase struct {
	Name string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Path string `json:"path"    dc:"角色ids"`
}

type PermissionGetListCommonReq struct {
	g.Meta `path:"/permission/list" method:"get" tags:"permission" summary:"permission list"`
	CommonPaginationReq
}

type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" pathription:"列表"`
	Page  int         `json:"page" pathription:"分页码"`
	Size  int         `json:"size" pathription:"分页数量"`
	Total int         `json:"total" pathription:"数据总数"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/permission/delete" method:"delete" tags:"permission" summary:"permission del"`
	Id     uint `v:"min:1#请选择需要删除的Permission" dc:"Permissionid"`
}

type PermissionDeleteRes struct{}
