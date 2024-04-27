package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsReq struct {
	g.Meta `path:"/goods/add" tags:"goods" method:"post" summary:"添加goods"`
	GoodsCommonAddUpdate
}

type GoodsCommonAddUpdate struct {
	PicUrl           string `json:"pic_url"         dc:"图片"`
	Name             string `json:"name"    v:"required#商品名称不能为空"         dc:"商品名称"`
	Price            int    `json:"price"    v:"required#商品名称不能为空"         dc:"价格 单位分"`
	Level1CategoryId int    `json:"level1_category_id" dc:"1级分类id"`
	Level2CategoryId int    `json:"level2_category_id" dc:"2级分类id"`
	Level3CategoryId int    `json:"level3_category_id" dc:"3级分类id"`
	Brand            string `json:"brand"            dc:"品牌"  v:"max-length:15"` //v: max only for int/float max-lenghth for string
	Stock            int    `json:"stock"            dc:"库存"`
	Sale             int    `json:"sale"             dc:"销量"`
	Tags             string `json:"tags"             dc:"标签"`
	DetailInfo       string `json:"detailInfo"       dc:"商品详情"`
}

type GoodsRes struct {
	Id uint `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"goods" summary:"删除goods接口"`
	Id     uint `v:"min:1#请选择需要删除的goods" dc:"goodsid"`
}
type GoodsDeleteRes struct{}

type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update/" method:"post" tags:"goods" summary:"修改goods接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的goods" dc:"goodsId"`
	GoodsCommonAddUpdate
}
type GoodsUpdateRes struct {
	Id uint `json:"id"`
}
type GoodsGetListCommonReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"goods" summary:"goods列表接口"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
