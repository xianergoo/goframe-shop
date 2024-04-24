package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CouponReq struct {
	g.Meta `path:"/coupon/add" tags:"coupon" method:"post" summary:"coupon add"`
	CouponCommonAddUpdate
}

type CouponCommonAddUpdate struct {
	Name       string `json:"name" v:"required#名称必填" dc:"分类名称"`
	Price      int    `json:"price"    v:"required#优惠券金额必填"      dc:"优惠券金额 分为单位 int类型"`
	GoodsIds   string `json:"goods_ids" dc:"关联使用的goods_ids  逗号分隔"`
	CategoryId uint8  `json:"category_id" dc:"关联使用的分类id"`
}
type CouponRes struct {
	CouponId uint `json:"couponId"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"coupon" summary:"coupon del"`
	Id     uint `v:"min:1#请选择需要删除的coupon" dc:"couponid"`
}

type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update/{Id}" method:"post" tags:"coupon" summary:"coupon modify"`
	Id     uint `json:"id"    v:"required#id不能为空"      dc:"coupon id"`
	CouponCommonAddUpdate
}
type CouponUpdateRes struct {
	Id uint `json:"id" `
}

type CouponGetListCommonReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"coupon" summary:"coupon list"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CouponGetListAllCommonReq struct {
	g.Meta `path:"/coupon/listall" method:"get" tags:"coupon" summary:"coupon list all"`
}
type CouponGetListAllCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
