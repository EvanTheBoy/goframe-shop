// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfo is the golang structure of table refund_info for DAO operations like Where/Data.
type RefundInfo struct {
	g.Meta    `orm:"table:refund_info, do:true"`
	Id        interface{} // 售后退款表
	Number    interface{} // 售后订单号
	OrderId   interface{} // 订单id
	GoodsId   interface{} // 要售后的商品id
	Reason    interface{} // 退款原因
	Status    interface{} // 状态 1待处理 2同意退款 3拒绝退款
	UserId    interface{} // 用户id
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
