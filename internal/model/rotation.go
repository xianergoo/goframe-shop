package model

import "github.com/gogf/gf/os/gtime"

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 内容模型
	Link   string // 栏目ID
	Sort   int    // 标题
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId uint `json:"rotation_id"`
}

// RotationUpdateInput 修改内容
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

// RotationGetList1Input 获取内容列表
type RotationGetList1Input struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetList1Output 查询列表结果
type RotationGetList1Output struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// RotationGetListInput 获取内容列表
type RotationGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationGetListOutput 查询列表结果
type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type RotationGetListOutputItem struct {
	Rotation *RotationListItem `json:"content"`
	// Category *RotationListCategoryItem `json:"category"`
	// User     *RotationListUserItem     `json:"user"`
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}

// RotationListItem ContentListItem 主要用于列表展示
type RotationListItem struct {
	Id        uint        `json:"id"`         // 自增ID
	PicUrl    string      `json:"pic_url"`    // 内容模型
	Link      string      `json:"link"`       // 栏目ID
	Sort      uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
