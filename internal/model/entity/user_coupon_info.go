// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	Id        int         `json:"id"        description:"用户优惠券表"`
	UserId    int         `json:"userId"    description:""`
	CouponId  int         `json:"couponId"  description:""`
	Status    int         `json:"status"    description:"状态：1可用 2已用 3过期"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
