package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Praise 承上启下
// Praise 内容管理
var Praise = cPraise{}

type cPraise struct{}

func (a *cPraise) Add(ctx context.Context, req *frontend.AddPraiseReq) (res *frontend.AddPraiseRes, err error) {

	data := model.AddPraiseInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddPraiseRes{Id: out.Id}, err
}

func (a *cPraise) Delete(ctx context.Context, req *frontend.DeletePraiseReq) (res *frontend.DeletePraiseRes, err error) {
	g.Dump(req)
	data := model.DeletePraiseInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	praise, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeletePraiseRes{Id: gconv.Uint(praise.Id)}, err
}

func (a *cPraise) GetList(ctx context.Context, req *frontend.ListPraiseReq) (res *frontend.ListPraiseRes, err error) {
	getListRes, err := service.Praise().GetList(ctx, model.PraiseListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	return &frontend.ListPraiseRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
