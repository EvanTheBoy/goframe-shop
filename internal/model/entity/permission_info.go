// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure for table permission_info.
type PermissionInfo struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"权限名称"`
	Path      string      `json:"path"      description:"路径"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
