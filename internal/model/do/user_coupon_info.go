// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure of table user_coupon_info for DAO operations like Where/Data.
type UserCouponInfo struct {
	g.Meta    `orm:"table:user_coupon_info, do:true"`
	Id        interface{} // 用户优惠券表
	UserId    interface{} //
	CouponId  interface{} //
	Status    interface{} // 状态：1可用 2已用 3过期
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
