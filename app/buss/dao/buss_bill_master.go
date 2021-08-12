// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/dao/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package dao

import (
    "gfast/app/buss/dao/internal"
    comModel "gfast/app/common/model"
    _ "github.com/gogf/gf/os/gtime"
)
// bussBillMasterDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type bussBillMasterDao struct {
	*internal.BussBillMasterDao
}
var (
	// BussBillMaster is globally public accessible object for table tools_gen_table operations.
	BussBillMaster bussBillMasterDao
)
func init() {
	BussBillMaster = bussBillMasterDao{
            internal.NewBussBillMasterDao(),
	}
}
// Fill with you ideas below.
// BussBillMasterSearchReq 分页请求参数
type BussBillMasterSearchReq struct {    
    BillId  string `p:"billId"` //主键    
    BillName  string `p:"billName"` //订单名称    
    BillType  string `p:"billType"` //订单类型    
    Status  string `p:"status"` //状态    
    BillStatus  string `p:"billStatus"` //订单状态    
    comModel.PageReq
}
// BussBillMasterAddReq 添加操作请求参数
type BussBillMasterAddReq struct {    
    AppId  int   `p:"appId" `    
    OwnerId  int64   `p:"ownerId" `    
    CustId  int64   `p:"custId" `    
    BillName  string   `p:"billName" v:"required#订单名称不能为空"`    
    BillType  string   `p:"billType" `    
    Status  uint   `p:"status" v:"required#状态不能为空"`    
    BillStatus  uint   `p:"billStatus" v:"required#订单状态不能为空"`    
    Remark  string   `p:"remark" `    
    CreateBy  uint   `p:"createBy" `    
    UpdateBy  uint   `p:"updateBy" `    
    BillSum  float64   `p:"billSum" `    
    BillNum  int   `p:"billNum" `    
}
// BussBillMasterEditReq 修改操作请求参数
type BussBillMasterEditReq struct {
    BillId    uint64  `p:"billId" v:"required#主键ID不能为空"`    
    AppId  int `p:"appId" `    
    OwnerId  int64 `p:"ownerId" `    
    CustId  int64 `p:"custId" `    
    BillName  string `p:"billName" v:"required#订单名称不能为空"`    
    BillType  string `p:"billType" `    
    Status  uint `p:"status" v:"required#状态不能为空"`    
    BillStatus  uint `p:"billStatus" v:"required#订单状态不能为空"`    
    Remark  string `p:"remark" `    
    CreateBy  uint `p:"createBy" `    
    UpdateBy  uint `p:"updateBy" `    
    BillSum  float64 `p:"billSum" `    
    BillNum  int `p:"billNum" `    
}
// BussBillMasterStatusReq 设置用户状态参数
type BussBillMasterStatusReq struct {
	BillId    uint64  `p:"billId" v:"required#主键ID不能为空"`
	Status uint   `p:"status" v:"required#状态不能为空"`
}
// BussBillMasterBillStatusReq 设置用户状态参数
type BussBillMasterBillStatusReq struct {
	BillId    uint64  `p:"billId" v:"required#主键ID不能为空"`
	BillStatus uint   `p:"billStatus" v:"required#订单状态不能为空"`
}
