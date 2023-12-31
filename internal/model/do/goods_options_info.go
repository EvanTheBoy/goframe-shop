// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptionsInfo is the golang structure of table goods_options_info for DAO operations like Where/Data.
type GoodsOptionsInfo struct {
	g.Meta    `orm:"table:goods_options_info, do:true"`
	Id        interface{} //
	GoodsId   interface{} // 商品id
	PicUrl    interface{} // 图片
	Name      interface{} // 商品名称
	Price     interface{} // 价格 单位分
	Stock     interface{} // 库存
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
