// Package api ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/api/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package api

import (
	"fmt"
	"gfast/app/buss/dao"
	"gfast/app/buss/service"
	sysApi "gfast/app/system/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)
type bussBillMaster struct {    
    sysApi.SystemBase    
}
var BussBillMaster = new(bussBillMaster)
// List 列表
func (c *bussBillMaster) List(r *ghttp.Request) {
	var req *dao.BussBillMasterSearchReq
	//获取参数
	fmt.Println("------------------------------------")
	fmt.Println(r)
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	total, page, list, err := service.BussBillMaster.GetList(req)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	result := g.Map{
		"currentPage": page,
		"total":       total,
		"list":        list,
	}
	c.SusJsonExit(r, result)
}
// Add 添加
func (c *bussBillMaster) Add(r *ghttp.Request) {
    var req *dao.BussBillMasterAddReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    
    err := service.BussBillMaster.Add(req)
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "添加成功")
}
// Get 获取
func (c *bussBillMaster) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.BussBillMaster.GetInfoById(id)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}
// Edit 修改
func (c *bussBillMaster) Edit(r *ghttp.Request) {
    var req *dao.BussBillMasterEditReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    
    err := service.BussBillMaster.Edit(req)
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "修改成功")
}
// Delete 删除
func (c *bussBillMaster) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.BussBillMaster.DeleteByIds(ids)
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}
// ChangeStatus 修改状态
func (c *bussBillMaster)ChangeStatus(r *ghttp.Request){
	   var req *dao.BussBillMasterStatusReq
	   //获取参数
	   if err := r.Parse(&req); err != nil {
	       c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	   }
	   if err := service.BussBillMaster.ChangeStatus(req); err != nil {
	       c.FailJsonExit(r, err.Error())
	   } else {
	       c.SusJsonExit(r, "状态设置成功")
	   }
}
// ChangeBillStatus 修改状态
func (c *bussBillMaster)ChangeBillStatus(r *ghttp.Request){
	   var req *dao.BussBillMasterBillStatusReq
	   //获取参数
	   if err := r.Parse(&req); err != nil {
	       c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	   }
	   if err := service.BussBillMaster.ChangeBillStatus(req); err != nil {
	       c.FailJsonExit(r, err.Error())
	   } else {
	       c.SusJsonExit(r, "状态设置成功")
	   }
}
