// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"图片名称"`
	Src       string      `json:"src"       description:"本地文件存储路径"`
	Url       string      `json:"url"       description:"URL地址"`
	UserId    int         `json:"userId"    description:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
