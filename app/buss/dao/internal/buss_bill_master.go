// ==========================================================================
// GFast自动生成dao internal操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/dao/internal/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package internal
import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)
// BussBillMasterDao is the manager for logic model data accessing and custom defined data operations functions management.
type BussBillMasterDao struct {
	gmvc.M                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      bussBillMasterColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB               // DB is the raw underlying database management object.
	Table  string               // Table is the underlying table name of the DAO.
}
// bussBillMasterColumns defines and stores column names for table buss_bill_master.
type bussBillMasterColumns struct {    
        BillId  string  // 主键    
        AppId  string  // 应用系统    
        OwnerId  string  // 拥有者    
        CustId  string  // 会员id    
        BillName  string  // 订单名称    
        BillType  string  // 订单类型    
        Status  string  // 状态    
        BillStatus  string  // 订单状态    
        Remark  string  // 备注    
        CreateBy  string  // 创建者    
        UpdateBy  string  // 更新者    
        CreatedAt  string  // 创建日期    
        UpdatedAt  string  // 修改日期    
        DeletedAt  string  // 删除日期    
        BillSum  string  // 订单合计    
        BillNum  string  // 订单数量    
}
// NewBussBillMasterDao creates and returns a new DAO object for table data access.
func NewBussBillMasterDao() *BussBillMasterDao {
	columns := bussBillMasterColumns{	    
            BillId:  "bill_id",        
            AppId:  "app_id",        
            OwnerId:  "owner_id",        
            CustId:  "cust_id",        
            BillName:  "bill_name",        
            BillType:  "bill_type",        
            Status:  "status",        
            BillStatus:  "bill_status",        
            Remark:  "remark",        
            CreateBy:  "create_by",        
            UpdateBy:  "update_by",        
            CreatedAt:  "created_at",        
            UpdatedAt:  "updated_at",        
            DeletedAt:  "deleted_at",        
            BillSum:  "bill_sum",        
            BillNum:  "bill_num",        
	}
	return &BussBillMasterDao{
		C:     columns,
		M:     g.DB("default").Model("buss_bill_master").Safe(),
		DB:    g.DB("default"),
		Table: "buss_bill_master",
	}
}
