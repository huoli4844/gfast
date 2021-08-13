// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"context"
	comModel "gfast/app/common/model"
	"gfast/app/system/model/internal"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData internal.SysDictData

// Fill with you ideas below.

// SelectDictPageReq 分页请求参数
type SelectDictPageReq struct {
	DictType  string `p:"dictType"`  //字典类型
	DictLabel string `p:"dictLabel"` //字典标签
	Status    string `p:"status"`    //状态
	comModel.PageReq
}

// GetDictReq 获取字典信息请求参数
type GetDictReq struct {
	DictType     string `p:"dictType" v:"required#字典类型不能为空"`
	DefaultValue string `p:"defaultValue"`
	Ctx          context.Context
}

// DictRes 完整的一个字典信息
type DictRes struct {
	Info   *DictTypeRes   `json:"info"`
	Values []*DictDataRes `json:"values"`
}

type DictTypeRes struct {
	DictName string `json:"name"`
	Remark   string `json:"remark"`
}

// DictDataRes 字典数据
type DictDataRes struct {
	DictValue string `json:"key"`
	DictLabel string `json:"value"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}

// DictDataAddReq 新增操作请求参数
type DictDataAddReq struct {
	DictLabel string `p:"dictLabel"  v:"required#字典标签不能为空"`
	DictValue string `p:"dictValue"  v:"required#字典键值不能为空"`
	DictType  string `p:"dictType"  v:"required#字典类型不能为空"`
	DictSort  int    `p:"dictSort"  v:"integer#排序只能为整数"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault int    `p:"isDefault" v:"required|in:0,1#系统默认不能为空|默认值只能为0或1"`
	Status    int    `p:"status"    v:"required|in:0,1#状态不能为空|状态只能为0或1"`
	Pic       string `p:"pic"`
	Remark    string `p:"remark"`
	CreateBy  uint64
}

// EditDictDataReq 修改字典数据操作请求参数
type EditDictDataReq struct {
	DictCode int `p:"dictCode" v:"required|min:1#主键ID不能为空|主键ID不能小于1"`
	UpdateBy uint64
	DictDataAddReq
}
