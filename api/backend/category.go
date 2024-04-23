package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CategoryReq struct {
	g.Meta `path:"/category/add" tags:"category" method:"post" summary:"category add"`
	CommonAddUpdate
}

type CommonAddUpdate struct {
	ParentId uint   `json:"parentId" dc:"父级分类ID"`
	Name     string `json:"name" v:"required#名称必填" dc:"分类名称"`
	PicUrl   string `json:"pic_url"    v:"required#图片链接不能为空"      dc:"图片"`
	Level    uint8  `json:"Level"    dc:"level"`
	Sort     uint8  `json:"sort" dc:"Sort"`
}
type CategoryRes struct {
	CategoryId uint `json:"categoryId"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"category" summary:"category del"`
	Id     uint `v:"min:1#请选择需要删除的category" dc:"categoryid"`
}

type CategoryDeleteRes struct{}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update/{Id}" method:"post" tags:"category" summary:"category modify"`
	Id     uint `json:"id"    v:"required#id不能为空"      dc:"category id"`
	CommonAddUpdate
}
type CategoryUpdateRes struct {
	Id uint `json:"id" `
}

type CategoryGetListCommonReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"category" summary:"category list"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
