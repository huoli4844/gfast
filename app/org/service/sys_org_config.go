// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-12 21:47:20
// 生成路径: gfast/app/org/service/sys_org_config.go
// 生成人：gfast
// ==========================================================================

package service

import (
    "context"    

	comModel "gfast/app/common/model"	

	"gfast/app/org/dao"
	"gfast/app/org/model"	

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type sysOrgConfig struct {
}

var SysOrgConfig = new(sysOrgConfig)

// GetList 获取任务列表
func (s *sysOrgConfig) GetList(req *dao.SysOrgConfigSearchReq,deptId int64) (total, page int, list []*model.SysOrgConfig, err error) {
	model := dao.SysOrgConfig.Ctx(req.Ctx)
	if req != nil {         

            if req.DeptId != "" {
                model = model.Where(dao.SysOrgConfig.Columns.DeptId+" = ?", req.DeptId)
            }        

            if req.ConfigName != "" {
                model = model.Where(dao.SysOrgConfig.Columns.ConfigName+" like ?", "%"+req.ConfigName+"%")
            }         

            if req.ConfigKey != "" {
                model = model.Where(dao.SysOrgConfig.Columns.ConfigKey+" = ?", req.ConfigKey)
            }         

            if req.AppId != "" {
                model = model.Where(dao.SysOrgConfig.Columns.AppId+" = ?", req.AppId)
            }        

            if req.AppName != "" {
                model = model.Where(dao.SysOrgConfig.Columns.AppName+" like ?", "%"+req.AppName+"%")
            }

            //如果不是平台超级管理员，那么只能看到自己公司的数据
            if deptId != 100 {
				model = model.Where(dao.SysOrgConfig.Columns.DeptId+" = ?", deptId)
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
    order:= "id asc"
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
func (s *sysOrgConfig) GetInfoById(id int64,ctx context.Context) (info *model.SysOrgConfig, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.SysOrgConfig.Ctx(ctx).Where(dao.SysOrgConfig.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *sysOrgConfig) Add(req *dao.SysOrgConfigAddReq,ctx context.Context) (err error) {
	_, err = dao.SysOrgConfig.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *sysOrgConfig) Edit(req *dao.SysOrgConfigEditReq,ctx context.Context) error {    

	_, err := dao.SysOrgConfig.Ctx(ctx).FieldsEx(dao.SysOrgConfig.Columns.Id,dao.SysOrgConfig.Columns.CreatedAt).Where(dao.SysOrgConfig.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *sysOrgConfig) DeleteByIds(ids []int,ctx context.Context) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}	

	_, err = dao.SysOrgConfig.Ctx(ctx).Delete(dao.SysOrgConfig.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

