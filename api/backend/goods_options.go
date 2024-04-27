package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsOptionsReq struct {
	g.Meta `path:"/goods/options/add" tags:"goods options" method:"post" summary:"添加goods"`
	GoodsOptionsCommonAddUpdate
}

type GoodsOptionsCommonAddUpdate struct {
	GoodsId uint   `json:"goods_id"        v:"required#主商品id不能为空" dc:"主商品id"`
	PicUrl  string `json:"pic_url"         dc:"图片"`
	Name    string `json:"name"            v:"required#商品名称不能为空"         dc:"商品名称"`
	Price   int    `json:"price"           v:"required#商品名称不能为空"         dc:"价格 单位分"`
	Stock   int    `json:"stock"            dc:"库存"`
	Sale    uint   `json:"sale"             dc:"销量"`
}

type GoodsOptionsRes struct {
	Id uint `json:"id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/goods/options/delete" method:"delete" tags:"goods options" summary:"goods options"`
	Id     uint `v:"min:1#请选择需要删除的goods" dc:"goods options id"`
}
type GoodsOptionsDeleteRes struct{}

type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/goods/options/update/" method:"post" tags:"goods options" summary:"goods options"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的goods options" dc:"goods options Id"`
	GoodsOptionsCommonAddUpdate
}
type GoodsOptionsUpdateRes struct {
	Id uint `json:"id"`
}
type GoodsOptionsGetListCommonReq struct {
	g.Meta `path:"/goods/options/list" method:"get" tags:"goods options" summary:"goods options列表接口"`
	CommonPaginationReq
}
type GoodsOptionsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
