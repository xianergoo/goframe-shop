package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"rotation" method:"post" summary:"rotation add"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}
type RotationRes struct {
	RotationId uint `json:"rotationId"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"rotation" summary:"rotation delete"`
	Id     uint `v:"min:1#请选择需要删除的rotation" dc:"rotationid"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/{Id}" method:"post" tags:"rotation" summary:"rotation modify"`
	Id     uint   `json:"id"    v:"required#图片id不能为空"      dc:"rotationid"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}
type RotationUpdateRes struct {
	Id uint `json:"id" `
}

type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" method:"get" tags:"rotation" summary:"rotation list"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
