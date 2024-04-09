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
	ILogin interface {
		// 获得头像上传路径
		GetAvatarUploadPath() string
		// 获得头像上传对应的URL前缀
		GetAvatarUploadUrlPrefix() string
		// 执行登录
		Login(ctx context.Context, in model.UserLoginInput) error
		// 注销
		Logout(ctx context.Context) error
		// 将密码按照内部算法进行加密
		EncryptPassword(passport, password string) string
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
