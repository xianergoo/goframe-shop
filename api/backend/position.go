package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PositionReq struct {
	g.Meta    `path:"/backend/position/add" tags:"position" method:"post" summary:"You first Position api"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link      string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空"      dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"    v:"required#商品ID不能为空"   dc:"商品Id"`
	Sort      int    `json:"sort"    dc:"排序"`
}
type PositionRes struct {
	PositionId uint `json:"positionId"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"position" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的position" dc:"positionid"`
}

type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" method:"post" tags:"position" summary:"修改position接口"`
	Id        uint   `json:"id"    v:"required#图片id不能为空"      dc:"positionid"`
	PicUrl    string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Link      string `json:"link"    v:"required#跳转链接不能为空"      dc:"跳转链接"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空"      dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"    v:"required#商品ID不能为空"   dc:"商品Id"`
	Sort      int    `json:"sort"    dc:"排序"`
}
type PositionUpdateRes struct {
	Id uint `json:"id" `
}

type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"position" summary:"轮播图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
