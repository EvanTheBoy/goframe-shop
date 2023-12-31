// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionInfo is the golang structure for table collection_info.
type CollectionInfo struct {
	Id        int         `json:"id"        description:""`
	UserId    int         `json:"userId"    description:"用户id"`
	ObjectId  int         `json:"objectId"  description:"对象id"`
	Type      int         `json:"type"      description:"收藏类型：1商品 2文章"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
