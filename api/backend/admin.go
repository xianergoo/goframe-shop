package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminReq struct {
	g.Meta   `path:"/admin/add" tags:"admin" method:"post" summary:"admin add"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级Admin"`
}

type AdminRes struct {
	AdminId uint `json:"admin_id"`
}
type AdminDeleteReq struct {
	g.Meta `path:"/admin/delete" method:"delete" tags:"admin" summary:"admin delete"`
	Id     uint `v:"min:1#请选择需要删除的Admin" dc:"Adminid"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/admin/update/{Id}" method:"post" tags:"admin" summary:"admin modify"`
	Id       uint   `json:"id"      v:"min:1#请选择需要修改的Admin" dc:"AdminId"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色ids"`
	IsAdmin  int    `json:"is_admin"    dc:"是否超级Admin"`
}
type AdminUpdateRes struct {
	Id uint `json:"id"`
}
type AdminGetListCommonReq struct {
	g.Meta `path:"/admin/list" method:"get" tags:"admin" summary:"admin list"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/admin/info" method:"get" tags:"admin" summary:"admin info"`
}

// for jwt
// type AdminGetInfoRes struct {
// 	Id          int    `json:"id"`
// 	IdentityKey string `json:"identity_key"`
// 	Payload     string `json:"payload"`
// }

// for gtoken
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}
