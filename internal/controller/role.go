package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

// Role 内容管理
var Role = cRole{}

type cRole struct{}

func (c *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.Id}, nil
}

// add permisson
func (c *cRole) AddPermission(ctx context.Context, req *backend.RoleAddPermssionReq) (res *backend.RoleAddPermssionRes, err error) {
	permission, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleAddPermssionRes{Id: permission.Id}, err
}

func (c *cRole) DeletePermission(ctx context.Context, req *backend.RoleDelPermssionReq) (res *backend.RoleDelPermssionRes, err error) {
	err = service.Role().DeletePermission(ctx, model.RoleDelPermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleDelPermssionRes{}, err
}

func (c *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (c *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return &backend.RoleUpdateRes{Id: req.Id}, err
}

func (c *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.RoleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}

func (c *cRole) GetInfo(ctx context.Context, req *backend.RoleInfoReq) (res *backend.RoleInfoRes, err error) {
	getInfoRes, err := service.Role().GetInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &backend.RoleInfoRes{
		Id:   getInfoRes.Id,
		Name: getInfoRes.Name,
		Desc: getInfoRes.Desc,
	}, nil
}
