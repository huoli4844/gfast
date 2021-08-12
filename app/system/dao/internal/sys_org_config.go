// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-12 20:12:41
// 生成路径: gfast/app/system/dao/internal/sys_org_config.go
// 生成人：gfast
// ==========================================================================

package internal

import (
    "context"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
)

// SysOrgConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysOrgConfigDao struct {
    Table   string         // Table is the underlying table name of the DAO.
    Group   string         // Group is the database configuration group name of current DAO.
    Columns SysOrgConfigColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SysOrgConfigColumns defines and stores column names for table sys_org_config.
type SysOrgConfigColumns struct {    

    Id  string  // 主键    

    DeptId  string  // 组织机构    

    ConfigId  string  // 参数主键    

    ConfigName  string  // 参数名称    

    ConfigKey  string  // 参数键名    

    ConfigValue  string  // 参数键值    

    ConfigType  string  // 系统内置    

    CreateBy  string  // 创建者    

    UpdateBy  string  // 更新者    

    Remark  string  // 备注    

    CreatedAt  string  // 创建时间    

    UpdatedAt  string  // 修改时间    

    DeletedAt  string  // 删除时间    

    AppId  string  // 系统类型    

    AppName  string  // 系统名称    

}

var sysOrgConfigColumns = SysOrgConfigColumns{    

    Id:  "id",    

    DeptId:  "dept_id",    

    ConfigId:  "config_id",    

    ConfigName:  "config_name",    

    ConfigKey:  "config_key",    

    ConfigValue:  "config_value",    

    ConfigType:  "config_type",    

    CreateBy:  "create_by",    

    UpdateBy:  "update_by",    

    Remark:  "remark",    

    CreatedAt:  "created_at",    

    UpdatedAt:  "updated_at",    

    DeletedAt:  "deleted_at",    

    AppId:  "app_id",    

    AppName:  "app_name",    

}

// NewSysOrgConfigDao creates and returns a new DAO object for table data access.
func NewSysOrgConfigDao() *SysOrgConfigDao {
	return &SysOrgConfigDao{
        Group:    "default",
        Table: "sys_org_config",
        Columns:sysOrgConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOrgConfigDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOrgConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOrgConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}