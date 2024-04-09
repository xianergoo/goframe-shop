package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
	res = &backend.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
