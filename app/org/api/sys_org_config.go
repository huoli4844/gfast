// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-12 21:47:20
// 生成路径: gfast/app/org/api/sys_org_config.go
// 生成人：gfast
// ==========================================================================

package api

import (    

    sysApi "gfast/app/system/api"    

    "gfast/app/org/dao"
    "gfast/app/org/service"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/util/gvalid"    

)

type sysOrgConfig struct {    

    sysApi.SystemBase    

}

var SysOrgConfig = new(sysOrgConfig)

// List 列表
func (c *sysOrgConfig) List(r *ghttp.Request) {
	var req *dao.SysOrgConfigSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.SysOrgConfig.GetList(req)
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
func (c *sysOrgConfig) Add(r *ghttp.Request) {
    var req *dao.SysOrgConfigAddReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    

    err := service.SysOrgConfig.Add(req,r.GetCtx())
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *sysOrgConfig) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.SysOrgConfig.GetInfoById(id,r.GetCtx())
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *sysOrgConfig) Edit(r *ghttp.Request) {
    var req *dao.SysOrgConfigEditReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    

    err := service.SysOrgConfig.Edit(req,r.GetCtx())
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *sysOrgConfig) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.SysOrgConfig.DeleteByIds(ids,r.GetCtx())
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

