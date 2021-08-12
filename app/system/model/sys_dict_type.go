// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/model/internal"
	"github.com/gogf/gf/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType internal.SysDictType

// Fill with you ideas below.

type ListSysDictTypeReq struct {
	DictName string `p:"dictName"` //字典名称
	DictType string `p:"dictType"` //字典类型
	Status   string `p:"status"`   //字典状态
	AppId   string `p:"appId"`   //系统类型
	comModel.PageReq
}

// SysDictTypeAddReq 新增操作请求参数
type SysDictTypeAddReq struct {
	DictName string `p:"dictName"  v:"required#字典名称不能为空"`
	DictType string `p:"dictType"  v:"required#字典类型不能为空"`
	Status   uint   `p:"status"  v:"required|in:0,1#状态不能为空|状态只能为0或1"`
	Remark   string `p:"remark"`
	AppId   string  `p:"appId"`   //系统类型
	CreateBy uint64
}

type SysDictTypeEditReq struct {
	SysDictTypeAddReq
	DictId   int64 `p:dictId v:required|min:1#主键ID不能为空|主键ID必须为大于0的值`
	UpdateBy uint64
}

type SysDictTypeInfoRes struct {
	DictId    uint64      `orm:"dict_id,primary"  json:"dictId"`    // 字典主键
	DictName  string      `orm:"dict_name"        json:"dictName"`  // 字典名称
	DictType  string      `orm:"dict_type,unique" json:"dictType"`  // 字典类型
	Status    uint        `orm:"status"           json:"status"`    // 状态（0正常 1停用）
	Remark    string      `orm:"remark"           json:"remark"`    // 备注
	CreatedAt *gtime.Time `orm:"created_at"       json:"createdAt"` // 创建日期
	AppId     string      `orm:"app_id"           json:"appId"`    // 系统类型
}
