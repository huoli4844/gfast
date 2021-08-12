// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/model/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package model
import "github.com/gogf/gf/os/gtime"
// BussBillMaster is the golang structure for table buss_bill_master.
type BussBillMaster struct {	
         BillId       uint64         `orm:"bill_id,primary" json:"billId"`    // 主键    
         AppId    int         `orm:"app_id" json:"appId"`    // 应用系统    
         OwnerId    int64         `orm:"owner_id" json:"ownerId"`    // 拥有者    
         CustId    int64         `orm:"cust_id" json:"custId"`    // 会员id    
         BillName    string         `orm:"bill_name" json:"billName"`    // 订单名称    
         BillType    string         `orm:"bill_type" json:"billType"`    // 订单类型    
         Status    uint         `orm:"status" json:"status"`    // 状态    
         BillStatus    uint         `orm:"bill_status" json:"billStatus"`    // 订单状态    
         Remark    string         `orm:"remark" json:"remark"`    // 备注    
         CreateBy    uint         `orm:"create_by" json:"createBy"`    // 创建者    
         UpdateBy    uint         `orm:"update_by" json:"updateBy"`    // 更新者    
         CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建日期    
         UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    // 修改日期    
         DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    // 删除日期    
         BillSum    float64         `orm:"bill_sum" json:"billSum"`    // 订单合计    
         BillNum    int         `orm:"bill_num" json:"billNum"`    // 订单数量    
}