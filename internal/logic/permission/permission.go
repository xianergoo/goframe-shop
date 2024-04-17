package permission

import (
	"context"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	//
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
			dao.PermissionInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PermissionInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.PermissionInfo.Columns().Id).
			Where(dao.PermissionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.PermissionInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
