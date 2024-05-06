package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) AddComment(ctx context.Context, in model.AddCommentInput) (out model.AddCommentOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	lastInsertID, err := dao.CommentInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddCommentOutput{Id: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sComment) DeleteComment(ctx context.Context, in model.DeleteCommentInput) (res *model.DeleteCommentOutput, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	}
	_, err = dao.CommentInfo.Ctx(ctx).Where(condition).Delete()
	if err != nil {
		return nil, err
	}
	return &model.DeleteCommentOutput{Id: in.Id}, err

}

// // Update 修改
// func (s *sComment) Update(ctx context.Context, in model.CommentUpdateInput) error {
// 	_, err := dao.CommentInfo.
// 		Ctx(ctx).
// 		Data(in).
// 		FieldsEx(dao.CommentInfo.Columns().Id).
// 		Where(dao.CommentInfo.Columns().Id, in.Id).
// 		Update()
// 	return err
// }

func (s *sComment) GetList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	// 分页查询
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.CommentInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	// filter
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
	}
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.CommentListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if in.Type == consts.CommentTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	}
	return

}
