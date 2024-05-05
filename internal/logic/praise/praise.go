package praise

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) AddPraise(ctx context.Context, in model.AddPraiseInput) (out model.AddPraiseOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.PraiseInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddPraiseOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sPraise) DeletePraise(ctx context.Context, in model.DeletePraiseInput) (res *model.DeletePraiseOutput, err error) {
	g.Dump(in)
	if in.Id != 0 {
		_, err = dao.PraiseInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeletePraiseOutput{Id: in.Id}, err
	} else {
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeletePraiseOutput{Id: gconv.Uint(id)}, err
	}
}

// // Update 修改
// func (s *sPraise) Update(ctx context.Context, in model.PraiseUpdateInput) error {
// 	_, err := dao.PraiseInfo.
// 		Ctx(ctx).
// 		Data(in).
// 		FieldsEx(dao.PraiseInfo.Columns().Id).
// 		Where(dao.PraiseInfo.Columns().Id, in.Id).
// 		Update()
// 	return err
// }

func (s *sPraise) GetList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	// 分页查询
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.PraiseInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.PraiseListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	// filter
	if in.Type != 0 {
		listModel = listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.PraiseListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	g.Dump(in)
	if in.Type == consts.PraiseTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	}
	return

}

func PraiseCount(ctx context.Context, objectId uint, praiseType uint) (count int, err error) {
	col := dao.PraiseInfo.Columns()
	condition := g.Map{
		col.ObjectId: objectId,
		col.Type:     praiseType,
	}
	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return count, err
}
