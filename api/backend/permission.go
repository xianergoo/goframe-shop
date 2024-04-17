package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionReq struct {
	g.Meta `path:"/backend/permission/add" method:"post" tags:"Permission" summary:"add permission"`
	Name   string `json:"name" v:"required"#require name" dc:"permission name"`
	Path   string `json:"path" path:"permission pathription"`
}

type PermissionRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update/{Id}" method:"post" tags:"Permission" summary:"修改permission接口"`
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
	g.Meta `path:"/backend/permission/list" method:"get" tags:"Permission" summary:"permission列表接口"`
	CommonPaginationReq
}

type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" pathription:"列表"`
	Page  int         `json:"page" pathription:"分页码"`
	Size  int         `json:"size" pathription:"分页数量"`
	Total int         `json:"total" pathription:"数据总数"`
}

type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete" method:"delete" tags:"Permission" summary:"删除Permission接口"`
	Id     uint `v:"min:1#请选择需要删除的Permission" dc:"Permissionid"`
}

type PermissionDeleteRes struct{}
