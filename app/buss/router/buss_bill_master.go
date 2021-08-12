// Package router ==========================================================================
// GFast自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-11 20:21:13
// 生成路径: gfast/app/buss/router/buss_bill_master.go
// 生成人：gfast
// ==========================================================================
package router

import (
    "gfast/app/buss/api"
    sysApi "gfast/app/system/api"
    "gfast/middleware"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
)
//加载路由
func init() {
    s := g.Server()
    s.Group("/", func(group *ghttp.RouterGroup) {
        group.Group("/buss", func(group *ghttp.RouterGroup) {
            group.Group("/bussBillMaster", func(group *ghttp.RouterGroup) {
                //gToken拦截器                
                sysApi.GfToken.Middleware(group)                
                //context拦截器
                group.Middleware(middleware.Ctx, middleware.Auth)                
                //后台操作日志记录
                group.Hook("/*", ghttp.HookAfterOutput, sysApi.SysOperLog.OperationLog)                
                group.GET("list", api.BussBillMaster.List)
                group.GET("get", api.BussBillMaster.Get)
                group.POST("add", api.BussBillMaster.Add)
                group.PUT("edit", api.BussBillMaster.Edit)
                group.DELETE("delete", api.BussBillMaster.Delete)                
                group.PUT("changeStatus",api.BussBillMaster.ChangeStatus)                
                group.PUT("changeBillStatus",api.BussBillMaster.ChangeBillStatus)                
            })
        })
    })
}
