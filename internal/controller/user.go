package controller

import (
	"context"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// User 内容管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *frontend.RegisterUserReq) (res *frontend.RegisterUserRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}

	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterUserRes{Id: out.Id}, nil
}

func (a *cUser) Info(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	res = &frontend.UserInfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserStatus))
	return res, nil

}

func (a *cUser) UpdatePassword(ctx context.Context, req *frontend.UpdatePasswordReq) (res *frontend.UpdatePasswordRes, err error) {
	//todo missing direct this api
	//log : 2024-04-28 23:32:04.930 [DEBU] {28e953056f7bca1744963e246f2d6168} [  2 ms] [default] [shop] [rows:1  ]
	//      SELECT `id`,`name`,`avatar`,`password`,`user_salt`,`sex`,`status`,`sign`,`secret_answer`,`created_at`,`updated_at`,`deleted_at` FROM `user_info` WHERE (`id`='8') AND `deleted_at` IS NULL LIMIT 1
	data := model.UpdatePasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.UpdatePasswordRes{Id: out.Id}, err

}
