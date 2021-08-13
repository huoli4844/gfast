// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-13 15:17:57
// 生成路径: gfast/app/system/dao/internal/demo_gen_other.go
// 生成人：gfast
// ==========================================================================

package internal

import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)

// DemoGenOtherDao is the manager for logic model data accessing and custom defined data operations functions management.
type DemoGenOtherDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns DemoGenOtherColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// DemoGenOtherColumns defines and stores column names for table demo_gen_other.
type DemoGenOtherColumns struct {    

    Id  string  // ID    

    Info  string  // 内容    

    Img  string  // 单图    

    Imgs  string  // 多图    

    File  string  // 单文件    

    Files  string  // 多文件    

    Remark  string  // 描述    

}

var demoGenOtherColumns = DemoGenOtherColumns{    

    Id:  "id",    

    Info:  "info",    

    Img:  "img",    

    Imgs:  "imgs",    

    File:  "file",    

    Files:  "files",    

    Remark:  "remark",    

}

// NewDemoGenOtherDao creates and returns a new DAO object for table data access.
func NewDemoGenOtherDao() *DemoGenOtherDao {
	return &DemoGenOtherDao{
        Group:    "default",
        Table: "demo_gen_other",
        Columns:demoGenOtherColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DemoGenOtherDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DemoGenOtherDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DemoGenOtherDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}