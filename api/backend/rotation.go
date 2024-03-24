package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/backend/rotation/add" tags:"rotation" method:"post" summary:"You first Rotation api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}
type RotationRes struct {
	RotationId uint `json:"rotationId"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/backend/rotation/delete" method:"delete" tags:"rotation" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的rotation" dc:"rotationid"`
}

type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/backend/rotation/update/{Id}" method:"post" tags:"rotation" summary:"修改rotation接口"`
	Id     uint   `json:"id"    v:"required#图片id不能为空"      dc:"rotationid"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link   string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}
type RotationUpdateRes struct {
	Id uint `json:"id" `
}

type RotationGetListReq struct {
	g.Meta `path:"/backend/rotation/list1" method:"get" tags:"rotation" summary:"rotation List"`
	// Page   int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	// Size   int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
	// Sort   int `json:"sort"    dc:"排序"`
	//CommonPaginationReq
}

type RotationGetListRes struct {
	// List  interface{} `json:"list" description:"列表"`
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type RotationGetListCommonReq struct {
	g.Meta `path:"/backend/rotation/list" method:"get" tags:"rotation" summary:"rotation List"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size   int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
	Sort   int `json:"sort"    dc:"排序"`
	//CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
