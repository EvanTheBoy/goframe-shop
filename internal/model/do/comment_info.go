// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure of table comment_info for DAO operations like Where/Data.
type CommentInfo struct {
	g.Meta    `orm:"table:comment_info, do:true"`
	Id        interface{} //
	ParentId  interface{} // 父级评论id
	UserId    interface{} //
	ObjectId  interface{} //
	Type      interface{} // 评论类型：1商品 2文章
	Content   interface{} // 评论内容
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
