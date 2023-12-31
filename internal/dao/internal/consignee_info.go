// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsigneeInfoDao is the data access object for table consignee_info.
type ConsigneeInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ConsigneeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ConsigneeInfoColumns defines and stores column names for table consignee_info.
type ConsigneeInfoColumns struct {
	Id        string // 收货地址表
	UserId    string //
	IsDefault string // 默认地址1  非默认0
	Name      string //
	Phone     string //
	Province  string //
	City      string //
	Town      string // 县区
	Street    string // 街道乡镇
	Detail    string // 地址详情
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// consigneeInfoColumns holds the columns for table consignee_info.
var consigneeInfoColumns = ConsigneeInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	IsDefault: "is_default",
	Name:      "name",
	Phone:     "phone",
	Province:  "province",
	City:      "city",
	Town:      "town",
	Street:    "street",
	Detail:    "detail",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewConsigneeInfoDao creates and returns a new DAO object for table data access.
func NewConsigneeInfoDao() *ConsigneeInfoDao {
	return &ConsigneeInfoDao{
		group:   "default",
		table:   "consignee_info",
		columns: consigneeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConsigneeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConsigneeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConsigneeInfoDao) Columns() ConsigneeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConsigneeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConsigneeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConsigneeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
