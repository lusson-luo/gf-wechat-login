// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WxUserDao is the data access object for table wx_user.
type WxUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns WxUserColumns // columns contains all the column names of Table for convenient usage.
}

// WxUserColumns defines and stores column names for table wx_user.
type WxUserColumns struct {
	UserId    string //
	OpenId    string //
	PhoneNo   string //
	AvatarUrl string //
	Nickname  string //
	Gender    string //
	CreateAt  string //
	UpdateAt  string //
}

// wxUserColumns holds the columns for table wx_user.
var wxUserColumns = WxUserColumns{
	UserId:    "user_id",
	OpenId:    "open_id",
	PhoneNo:   "phone_no",
	AvatarUrl: "avatar_url",
	Nickname:  "nickname",
	Gender:    "gender",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// NewWxUserDao creates and returns a new DAO object for table data access.
func NewWxUserDao() *WxUserDao {
	return &WxUserDao{
		group:   "default",
		table:   "wx_user",
		columns: wxUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WxUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WxUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WxUserDao) Columns() WxUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WxUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WxUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WxUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
