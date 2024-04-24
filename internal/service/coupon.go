// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop/internal/model"
)

type (
	ICoupon interface {
		Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.CouponUpdateInput) error
		GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error)
		// GetList 查询内容列表
		GetListAll(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error)
	}
)

var (
	localCoupon ICoupon
)

func Coupon() ICoupon {
	if localCoupon == nil {
		panic("implement not found for interface ICoupon, forgot register?")
	}
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
