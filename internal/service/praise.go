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
	IPraise interface {
		AddPraise(ctx context.Context, in model.AddPraiseInput) (out model.AddPraiseOutput, err error)
		// Delete 删除
		DeletePraise(ctx context.Context, in model.DeletePraiseInput) (res *model.DeletePraiseOutput, err error)
		GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error)
	}
)

var (
	localPraise IPraise
)

func Praise() IPraise {
	if localPraise == nil {
		panic("implement not found for interface IPraise, forgot register?")
	}
	return localPraise
}

func RegisterPraise(i IPraise) {
	localPraise = i
}
