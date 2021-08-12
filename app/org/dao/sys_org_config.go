// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2021-08-12 21:47:20
// 生成路径: gfast/app/org/dao/sys_org_config.go
// 生成人：gfast
// ==========================================================================

package dao

import (
    comModel "gfast/app/common/model"
    "gfast/app/org/dao/internal"

    _ "github.com/gogf/gf/os/gtime"
)

// sysOrgConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysOrgConfigDao struct {
	*internal.SysOrgConfigDao
}

var (
    // SysOrgConfig is globally public accessible object for table tools_gen_table operations.
    SysOrgConfig = sysOrgConfigDao{
        internal.NewSysOrgConfigDao(),
    }
)

// Fill with you ideas below.

// SysOrgConfigSearchReq 分页请求参数
type SysOrgConfigSearchReq struct {    

    DeptId  string `p:"deptId"` //组织机构id 就是部门ID    

    ConfigName  string `p:"configName"` //参数名称    

    ConfigKey  string `p:"configKey"` //参数键名    

    AppId  string `p:"appId"` //系统类型    

    AppName  string `p:"appName"` //系统名称    

    comModel.PageReq
}

// SysOrgConfigAddReq 添加操作请求参数
type SysOrgConfigAddReq struct {    

    DeptId  int64   `p:"deptId" `    

    ConfigId  int   `p:"configId" `    

    ConfigName  string   `p:"configName" v:"required#参数名称不能为空"`    

    ConfigKey  string   `p:"configKey" `    

    ConfigValue  string   `p:"configValue" `    

    ConfigType  int   `p:"configType" `    

    CreateBy  uint   `p:"createBy" `    

    UpdateBy  uint   `p:"updateBy" `    

    Remark  string   `p:"remark" `    

    AppId  int   `p:"appId" `    

    AppName  string   `p:"appName" v:"required#系统名称不能为空"`    

}

// SysOrgConfigEditReq 修改操作请求参数
type SysOrgConfigEditReq struct {
    Id    int64  `p:"id" v:"required#主键ID不能为空"`    

    ConfigValue  string `p:"configValue" `    

    Remark  string `p:"remark" `    

}

