// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptionsInfo is the golang structure for table goods_options_info.
type GoodsOptionsInfo struct {
	Id        int         `json:"id"        description:""`
	GoodsId   int         `json:"goodsId"   description:"商品id"`
	PicUrl    string      `json:"picUrl"    description:"图片"`
	Name      string      `json:"name"      description:"商品名称"`
	Price     int         `json:"price"     description:"价格 单位分"`
	Stock     int         `json:"stock"     description:"库存"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
