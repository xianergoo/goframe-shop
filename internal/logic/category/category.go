package category

import (
	"context"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	glog.Info(ctx, in)
	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sCategory) Delete(ctx context.Context, id uint) error {
	_, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Id: id,
	}).Delete()
	return err

}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) error {
	_, err := dao.CategoryInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return err
}

func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	// 分页查询
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.CategoryInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.CategoryGetListOutput{
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
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.CategoryGetListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return

}

// GetList 查询内容列表
func (s *sCategory) GetListAll(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {

	//查询全部
	m := dao.CategoryInfo.Ctx(ctx)
	out = &model.CategoryGetListOutput{}

	listModel := m
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	/*
		//
		var list []*entity.CategoryInfo
		if err := listModel.Scan(&list); err != nil {
			return out, err
		}

		if len(list) == 0 {
			return out, nil
		} */
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return

}
