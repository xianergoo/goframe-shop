package controller

import (
	"context"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// 承上启下
// Goods内容管理
var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
	data := model.GoodsCreateInput{}
	gconv.Scan(req, &data)
	glog.Info(ctx, data)
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsRes{Id: out.Id}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	data := model.GoodsUpdateInput{}
	gconv.Struct(req, &data)
	err = service.Goods().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsUpdateRes{Id: req.Id}, nil
}

func (a *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.GoodsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
