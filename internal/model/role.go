package model

// RoleCreateUpdateBase 创建/修改内容基类
type RoleCreateUpdateBase struct {
	Name string // 内容模型
	Desc string
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建内容返回结果
type RoleCreateOutput struct {
	Id uint `json:"role_id"`
}

// RoleUpdateInput 修改内容
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint
}

// RoleGetListInput 获取内容列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RoleGetListOutput 查询列表结果
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type RoleGetListOutputItem struct {
	Id   uint   `json:"id"` // 自增ID
	Name string `json:"name"    description:"用户名"`
}
