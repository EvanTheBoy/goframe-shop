// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfo is the golang structure for table refund_info.
type RefundInfo struct {
	Id        int         `json:"id"        description:"售后退款表"`
	Number    string      `json:"number"    description:"售后订单号"`
	OrderId   int         `json:"orderId"   description:"订单id"`
	GoodsId   int         `json:"goodsId"   description:"要售后的商品id"`
	Reason    string      `json:"reason"    description:"退款原因"`
	Status    int         `json:"status"    description:"状态 1待处理 2同意退款 3拒绝退款"`
	UserId    int         `json:"userId"    description:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
