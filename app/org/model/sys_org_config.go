// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-12 21:47:20
// 生成路径: gfast/app/org/model/sys_org_config.go
// 生成人：gfast
// ==========================================================================

package model

import "github.com/gogf/gf/os/gtime"

// SysOrgConfig is the golang structure for table sys_org_config.
type SysOrgConfig struct {	

         Id       int64         `orm:"id,primary" json:"id"`    // 主键    

         DeptId    int64         `orm:"dept_id" json:"deptId"`    // 组织机构id 就是部门ID    

         ConfigId    int         `orm:"config_id" json:"configId"`    // 参数主键    

         ConfigName    string         `orm:"config_name" json:"configName"`    // 参数名称    

         ConfigKey    string         `orm:"config_key" json:"configKey"`    // 参数键名    

         ConfigValue    string         `orm:"config_value" json:"configValue"`    // 参数键值    

         ConfigType    int         `orm:"config_type" json:"configType"`    // 系统内置（Y是 N否）    

         CreateBy    uint         `orm:"create_by" json:"createBy"`    // 创建者    

         UpdateBy    uint         `orm:"update_by" json:"updateBy"`    // 更新者    

         Remark    string         `orm:"remark" json:"remark"`    // 备注    

         CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    

         UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    // 修改时间    

         DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    // 删除时间    

         AppId    int         `orm:"app_id" json:"appId"`    // 系统类型    

         AppName    string         `orm:"app_name" json:"appName"`    // 系统名称    

}