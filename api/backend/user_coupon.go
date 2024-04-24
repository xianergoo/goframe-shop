package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponReq struct {
	g.Meta `path:"/user/coupon/add" tags:"coupon" method:"post" summary:"添加coupon"`
	UserCouponCommonAddUpdate
}

type UserCouponCommonAddUpdate struct {
	UserId   uint  `json:"user_id" v:"required#用户id必填" dc:"用户id"`
	CouponId uint  `json:"coupon_id" v:"required#优惠券id必填" dc:"可用的商品分类"`
	Status   uint8 `json:"status" dc:"状态"`
}

type UserCouponRes struct {
	Id uint `json:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" method:"delete" tags:"coupon" summary:"删除coupon接口"`
	Id     uint `v:"min:1#请选择需要删除的coupon" dc:"couponid"`
}
type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update/" method:"post" tags:"coupon" summary:"修改coupon接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的coupon" dc:"couponId"`
	UserCouponCommonAddUpdate
}
type UserCouponUpdateRes struct {
	Id uint `json:"id"`
}
type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user/coupon/list" method:"get" tags:"coupon" summary:"coupon列表接口"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
