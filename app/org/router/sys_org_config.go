// ==========================================================================
// GFast自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-12 21:47:20
// 生成路径: gfast/app/org/router/sys_org_config.go
// 生成人：gfast
// ==========================================================================
package router

import (
    "gfast/app/org/api"
    "gfast/middleware"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"    

    sysApi "gfast/app/system/api"    

)

//加载路由
func init() {
    s := g.Server()
    s.Group("/", func(group *ghttp.RouterGroup) {
        group.Group("/org", func(group *ghttp.RouterGroup) {
            group.Group("/sysOrgConfig", func(group *ghttp.RouterGroup) {
                //gToken拦截器                

                sysApi.GfToken.Middleware(group)                

                //context拦截器
                group.Middleware(middleware.Ctx, middleware.Auth)                

                //后台操作日志记录
                group.Hook("/*", ghttp.HookAfterOutput, sysApi.SysOperLog.OperationLog)                

                group.GET("list", api.SysOrgConfig.List)
                group.GET("get", api.SysOrgConfig.Get)
                group.POST("add", api.SysOrgConfig.Add)
                group.PUT("edit", api.SysOrgConfig.Edit)
                group.DELETE("delete", api.SysOrgConfig.Delete)                

            })
        })
    })
}
