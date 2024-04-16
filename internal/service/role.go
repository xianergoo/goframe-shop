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
	IRole interface {
		Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error)
		// add permission
		AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Delete 删除 permission
		DeletePermission(ctx context.Context, in model.RoleDelPermissionInput) error
		// Update 修改
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
		// GetList 查询内容列表
		GetInfo(ctx context.Context, id uint) (out *model.RoleGetListOutputItem, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
