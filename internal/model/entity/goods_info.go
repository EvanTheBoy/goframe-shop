// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	Id               int         `json:"id"               description:""`
	PicUrl           string      `json:"picUrl"           description:"图片"`
	Name             string      `json:"name"             description:"商品名称"`
	Price            int         `json:"price"            description:"价格 单位分"`
	Level1CategoryId int         `json:"level1CategoryId" description:"1级分类id"`
	Level2CategoryId int         `json:"level2CategoryId" description:"2级分类id"`
	Level3CategoryId int         `json:"level3CategoryId" description:"3级分类id"`
	Brand            string      `json:"brand"            description:"品牌"`
	Stock            int         `json:"stock"            description:"库存"`
	Sale             int         `json:"sale"             description:"销量"`
	Tags             string      `json:"tags"             description:"标签"`
	DetailInfo       string      `json:"detailInfo"       description:"商品详情"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:""`
	DeletedAt        *gtime.Time `json:"deletedAt"        description:""`
}
