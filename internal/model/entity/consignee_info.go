// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure for table consignee_info.
type ConsigneeInfo struct {
	Id        int         `json:"id"        description:"收货地址表"`
	UserId    int         `json:"userId"    description:""`
	IsDefault int         `json:"isDefault" description:"默认地址1  非默认0"`
	Name      string      `json:"name"      description:""`
	Phone     string      `json:"phone"     description:""`
	Province  string      `json:"province"  description:""`
	City      string      `json:"city"      description:""`
	Town      string      `json:"town"      description:"县区"`
	Street    string      `json:"street"    description:"街道乡镇"`
	Detail    string      `json:"detail"    description:"地址详情"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
