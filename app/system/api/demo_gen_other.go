// ==========================================================================
// GFast自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-13 15:17:57
// 生成路径: gfast/app/system/api/demo_gen_other.go
// 生成人：gfast
// ==========================================================================

package api

import (    

    "gfast/library"    

    "gfast/app/system/dao"
    "gfast/app/system/service"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/util/gvalid"    

    "github.com/gogf/gf/encoding/gjson"    

)

type demoGenOther struct {    

    SystemBase    

}

var DemoGenOther = new(demoGenOther)

// List 列表
func (c *demoGenOther) List(r *ghttp.Request) {
	var req *dao.DemoGenOtherSearchReq
	//获取参数
	if err := r.Parse(&req); err != nil {
		c.FailJsonExit(r, err.(gvalid.Error).FirstString())
	}
	req.Ctx = r.GetCtx()
	total, page, list, err := service.DemoGenOther.GetList(req)
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
func (c *demoGenOther) Add(r *ghttp.Request) {
    var req *dao.DemoGenOtherAddReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    

    upImgs:=gjson.New(req.Imgs)
    for _,obj:=range upImgs.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.Imgs = upImgs.MustToJsonString()    

    upFile:=gjson.New(req.File)
    for _,obj:=range upFile.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.File = upFile.MustToJsonString()    

    upFiles:=gjson.New(req.Files)
    for _,obj:=range upFiles.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.Files = upFiles.MustToJsonString()    

    err := service.DemoGenOther.Add(req,r.GetCtx())
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "添加成功")
}

// Get 获取
func (c *demoGenOther) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	info, err := service.DemoGenOther.GetInfoById(id,r.GetCtx())
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, info)
}

// Edit 修改
func (c *demoGenOther) Edit(r *ghttp.Request) {
    var req *dao.DemoGenOtherEditReq
    //获取参数
    if err := r.Parse(&req); err != nil {
        c.FailJsonExit(r, err.(gvalid.Error).FirstString())
    }    

    upImgs:=gjson.New(req.Imgs)
    for _,obj:=range upImgs.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.Imgs = upImgs.MustToJsonString()    

    upFile:=gjson.New(req.File)
    for _,obj:=range upFile.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.File = upFile.MustToJsonString()    

    upFiles:=gjson.New(req.Files)
    for _,obj:=range upFiles.Array(){
        mp := obj.(g.MapStrAny)
        var err error
        mp["url"],err = library.GetFilesPath(mp["url"].(string))
        if err!=nil{
            c.FailJsonExit(r, err.Error())
        }
    }
    req.Files = upFiles.MustToJsonString()    

    err := service.DemoGenOther.Edit(req,r.GetCtx())
    if err != nil {
        c.FailJsonExit(r, err.Error())
    }
    c.SusJsonExit(r, "修改成功")
}

// Delete 删除
func (c *demoGenOther) Delete(r *ghttp.Request) {
	ids := r.GetInts("ids")
	err := service.DemoGenOther.DeleteByIds(ids,r.GetCtx())
	if err != nil {
		c.FailJsonExit(r, err.Error())
	}
	c.SusJsonExit(r, "删除成功")
}

