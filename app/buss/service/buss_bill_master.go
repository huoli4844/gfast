// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/service/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package service
import (    
	comModel "gfast/app/common/model"	
	"gfast/app/buss/dao"
	"gfast/app/buss/model"	
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)
type bussBillMaster struct {
}
var BussBillMaster = new(bussBillMaster)
// GetList 获取任务列表
func (s *bussBillMaster) GetList(req *dao.BussBillMasterSearchReq) (total, page int, list []*model.BussBillMaster, err error) {
	model := dao.BussBillMaster.M
	if req != nil {         
            if req.BillId != "" {
                model = model.Where(dao.BussBillMaster.C.BillId+" = ?", req.BillId)
            }        
            if req.BillName != "" {
                model = model.Where(dao.BussBillMaster.C.BillName+" like ?", "%"+req.BillName+"%")
            }         
            if req.BillType != "" {
                model = model.Where(dao.BussBillMaster.C.BillType+" = ?", req.BillType)
            }         
            if req.Status != "" {
                model = model.Where(dao.BussBillMaster.C.Status+" = ?", req.Status)
            }         
            if req.BillStatus != "" {
                model = model.Where(dao.BussBillMaster.C.BillStatus+" = ?", req.BillStatus)
            }            
	}
	total, err = model.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}    
    if req.PageNum == 0 {
        req.PageNum = 1
    }
    page = req.PageNum
    if req.PageSize == 0 {
        req.PageSize = comModel.PageSize
    }
    order:= "bill_id asc"
    if req.OrderBy!=""{
        order = req.OrderBy
    }
    err = model.Page(page, req.PageSize).Order(order).Scan(&list)    
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
	}
	return
}
// GetInfoById 通过id获取
func (s *bussBillMaster) GetInfoById(id int64) (info *model.BussBillMaster, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.BussBillMaster.Where(dao.BussBillMaster.C.BillId, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}
// Add 添加
func (s *bussBillMaster) Add(req *dao.BussBillMasterAddReq) (err error) {
	_, err = dao.BussBillMaster.Insert(req)
	return
}
// Edit 修改
func (s *bussBillMaster) Edit(req *dao.BussBillMasterEditReq) error {    
	_, err := dao.BussBillMaster.FieldsEx(dao.BussBillMaster.C.BillId,dao.BussBillMaster.C.CreatedAt).Where(dao.BussBillMaster.C.BillId, req.BillId).
		Update(req)
	return err
}
// DeleteByIds 删除
func (s *bussBillMaster) DeleteByIds(ids []int) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
	}	
	_, err = dao.BussBillMaster.Delete(dao.BussBillMaster.C.BillId+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}
// ChangeStatus 修改状态
func (s *bussBillMaster) ChangeStatus(req *dao.BussBillMasterStatusReq) error {
	_, err := dao.BussBillMaster.WherePri(req.BillId).Update(g.Map{
		dao.BussBillMaster.C.Status: req.Status,
	})
	return err
}
// ChangeBillStatus 修改状态
func (s *bussBillMaster) ChangeBillStatus(req *dao.BussBillMasterBillStatusReq) error {
	_, err := dao.BussBillMaster.WherePri(req.BillId).Update(g.Map{
		dao.BussBillMaster.C.BillStatus: req.BillStatus,
	})
	return err
}
