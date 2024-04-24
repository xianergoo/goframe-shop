package model

import "github.com/gogf/gf/os/gtime"

// CouponCreateUpdateBase 创建/修改内容基类
type CouponCreateUpdateBase struct {
	Name       string
	Price      int
	GoodsIds   string
	CategoryId uint8
}

// CouponCreateInput 创建内容
type CouponCreateInput struct {
	CouponCreateUpdateBase
}

// CouponCreateOutput 创建内容返回结果
type CouponCreateOutput struct {
	CouponId uint `json:"rotation_id"`
}

// CouponUpdateInput 修改内容
type CouponUpdateInput struct {
	CouponCreateUpdateBase
	Id uint
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

type CouponGetListOutputItem struct {
	Id         uint        `json:"id"`
	Name       string      `json:"name"`
	Price      int         `json:"price"`
	GoodsIds   string      `json:"level"`      //级别
	CategoryId uint        `json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedAt  *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"` // 修改时间
}
