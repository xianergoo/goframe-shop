package role

import (
	"context"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/database/gdb"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	//
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{Id: uint(lastInsertID)}, err
}

// add permission
func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error) {
	//
	id, err := dao.RolePermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleAddPermissionOutput{Id: uint(id)}, err
}

// Delete 删除
func (s *sRole) Delete(ctx context.Context, id uint) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
			dao.RoleInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Delete 删除 permission
func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDelPermissionInput) error {
	db := dao.RolePermissionInfo.Ctx(ctx)
	_, err := db.Where(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	return err
}

// Update 修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.RoleInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RoleInfo.Columns().Id).
			Where(dao.RoleInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.RoleInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.RoleGetListOutput{
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

// GetList 查询内容列表
func (s *sRole) GetInfo(ctx context.Context, id uint) (out *model.RoleGetListOutputItem, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.RoleInfo.Ctx(ctx)

	//4. 再查询count，判断有无数据
	out = &model.RoleGetListOutputItem{}
	m.Where("id", id).Scan(&out)
	return
}
