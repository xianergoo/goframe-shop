package user

import (
	"context"
	"errors"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/service"
	"goframe-shop/utility"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out *model.UpdatePasswordOutput, err error) {
	glog.Debug(ctx, in)
	if in.Password != "" {
		UserSalt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, UserSalt)
		in.UserSalt = UserSalt
	}
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(userInfo)
	if err != nil {
		return &model.UpdatePasswordOutput{}, errors.New(consts.ErrLoginFaulMsg)
	}
	if userInfo.SecretAnswer != in.SecretAnswer {
		return &model.UpdatePasswordOutput{}, errors.New(consts.ErrSecretAnswerMsg)
	}

	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return &model.UpdatePasswordOutput{}, err
	}
	return &model.UpdatePasswordOutput{Id: userId}, err
}
