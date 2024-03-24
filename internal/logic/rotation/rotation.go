package Rotation

import (
	"context"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		// Unscoped().Delete() 删除数据
		// Delete()  更新Delete time 字段
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Unscoped().Delete() //
		return err
	})

}

// Delete1 删除
func (s *sRotation) Delete1(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		// Unscoped().Delete() 删除数据
		// Delete()  更新Delete time 字段
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Unscoped().Delete() //
		return err
	})

}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

func (s *sRotation) GetList1(ctx context.Context, in model.RotationGetList1Input) (out model.RotationGetList1Output, err error) {
	out = model.RotationGetList1Output{
		Page: in.Page,
		Size: in.Size,
	}
	return
}

// GetList 查询内容列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	var (
		m = dao.RotationInfo.Ctx(ctx)
	)
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 默认查询topic
	// if in.Type != "" {
	// 	m = m.Where(dao.RotationInfo.Columns().Type, in.Type)
	// } else {
	// 	m = m.Where(dao.RotationInfo.Columns().Type, consts.RotationTypeTopic)
	// }
	// 栏目检索
	// if in.CategoryId > 0 {
	// 	idArray, err := service.Category().GetSubIdList(ctx, in.CategoryId)
	// 	if err != nil {
	// 		return out, err
	// 	}
	// 	m = m.Where(dao.Rotation.Columns().CategoryId, idArray)
	// }
	// 管理员可以查看所有文章
	// if in.UserId > 0 && !service.User().IsAdmin(ctx, in.UserId) {
	// 	m = m.Where(dao.Rotation.Columns().UserId, in.UserId)
	// }
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	switch in.Sort {
	case consts.RotationSortActive:
		listModel = listModel.OrderDesc(dao.RotationInfo.Columns().UpdatedAt)

	default:
		listModel = listModel.OrderDesc(dao.RotationInfo.Columns().Sort)
	}
	// 执行查询
	var list []*entity.RotationInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Rotation
	if err := listModel.ScanList(&out.List, "Rotation"); err != nil {
		return out, err
	}
	return
}
