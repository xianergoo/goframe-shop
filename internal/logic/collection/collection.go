package collection

import (
	"context"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (out model.AddCollectionOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.CollectionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddCollectionOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (res *model.DeleteCollectionOutput, err error) {
	if in.Id != 0 {
		_, err = dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeleteCollectionOutput{Id: in.Id}, err
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).Where(in).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeleteCollectionOutput{Id: gconv.Uint(id)}, err
	}
}

// // Update 修改
// func (s *sCollection) Update(ctx context.Context, in model.CollectionUpdateInput) error {
// 	_, err := dao.CollectionInfo.
// 		Ctx(ctx).
// 		Data(in).
// 		FieldsEx(dao.CollectionInfo.Columns().Id).
// 		Where(dao.CollectionInfo.Columns().Id, in.Id).
// 		Update()
// 	return err
// }

func (s *sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	// 分页查询
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.CollectionInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.CollectionListOutput{
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
	out.List = make([]model.CollectionListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.WithAll().Scan(&out.List); err != nil {
		return out, err
	}
	return

}
