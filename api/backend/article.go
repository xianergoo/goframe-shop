package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"article" method:"post" summary:"添加article"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	Title   string `json:"title"         dc:"title" v:"required#title is required"`
	Desc    string `json:"desc"           dc:"Description"`
	PicUrl  string `json:"pic_url"            dc:"picture url"  ` //v: max only for int/float max-lenghth for string
	IsAdmin uint   `json:"is_admin"            dc:"is admin"`
	Praise  int    `json:"praise"    dc:"点赞数"`
	Detail  string `json:"detail"    dc:"文章详情"`
}

type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"article" summary:"删除article接口"`
	Id     uint `v:"min:1#请选择需要删除的article" dc:"articleid"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update/" method:"post" tags:"article" summary:"修改article接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的article" dc:"articleId"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}
type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"article" summary:"article列表接口"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
