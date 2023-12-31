// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentInfoDao is the data access object for table comment_info.
type CommentInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CommentInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CommentInfoColumns defines and stores column names for table comment_info.
type CommentInfoColumns struct {
	Id        string //
	ParentId  string // 父级评论id
	UserId    string //
	ObjectId  string //
	Type      string // 评论类型：1商品 2文章
	Content   string // 评论内容
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// commentInfoColumns holds the columns for table comment_info.
var commentInfoColumns = CommentInfoColumns{
	Id:        "id",
	ParentId:  "parent_id",
	UserId:    "user_id",
	ObjectId:  "object_id",
	Type:      "type",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewCommentInfoDao creates and returns a new DAO object for table data access.
func NewCommentInfoDao() *CommentInfoDao {
	return &CommentInfoDao{
		group:   "default",
		table:   "comment_info",
		columns: commentInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommentInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommentInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommentInfoDao) Columns() CommentInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommentInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommentInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommentInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
