package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"admin" method:"post" summary:"You first Admin api"`
	Name     string `json:"name"     v:"required#Name can not be empty" description:"用户名"`
	Password string `json:"password" v:"required#Password can not be empty" description:"密码"`
	RoleIds  string `json:"roleIds"   description:"角色ids"`
	IsAdmin  int    `json:"isAdmin"   description:"是否超级管理员"`
}

type AdminRes struct {
	Id uint `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"admin" summary:"删除admin接口"`
	Id     uint `v:"min:1#请选择需要删除的admin" dc:"admin_id"`
}

type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"admin" summary:"修改admin接口"`
	Id       uint   `json:"id"    v:"required#id不能为空"      dc:"admin_id"`
	Name     string `json:"name"     v:"required#Name can not be empty" description:"用户名"`
	Password string `json:"password" v:"required#Password can not be empty" description:"密码"`
	RoleIds  string `json:"roleIds"   description:"角色ids"`
	IsAdmin  int    `json:"isAdmin"   description:"是否超级管理员"`
}
type AdminUpdateRes struct {
	Id uint `json:"id" `
}

type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" method:"get" tags:"admin" summary:"admin列表接口"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
