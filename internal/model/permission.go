package model

// PermissionCreateUpdateBase 创建/修改内容基类
type PermissionCreateUpdateBase struct {
	Name string // 内容模型
	Path string
}

// PermissionCreateInput 创建内容
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建内容返回结果
type PermissionCreateOutput struct {
	Id uint `json:"permission_id"`
}

// PermissionUpdateInput 修改内容
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint
}

// PermissionGetListInput 获取内容列表
type PermissionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PermissionGetListOutput 查询列表结果
type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" pathription:"列表"`
	Page  int                           `json:"page" pathription:"分页码"`
	Size  int                           `json:"size" pathription:"分页数量"`
	Total int                           `json:"total" pathription:"数据总数"`
}

type PermissionGetListOutputItem struct {
	Id   uint   `json:"id"` // 自增ID
	Name string `json:"name"    pathription:"用户名"`
	Path string `json:"path"`
}
