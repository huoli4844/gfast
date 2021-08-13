// ==========================================================================
// GFast自动生成业务逻辑层相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2021-08-13 15:17:57
// 生成路径: gfast/app/system/service/demo_gen_other.go
// 生成人：gfast
// ==========================================================================

package service

import (
    "context"    

	comModel "gfast/app/common/model"	

	"gfast/app/system/dao"
	"gfast/app/system/model"	

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type demoGenOther struct {
}

var DemoGenOther = new(demoGenOther)

// GetList 获取任务列表
func (s *demoGenOther) GetList(req *dao.DemoGenOtherSearchReq) (total, page int, list []*model.DemoGenOther, err error) {
	model := dao.DemoGenOther.Ctx(req.Ctx)
	if req != nil {         

            if req.Info != "" {
                model = model.Where(dao.DemoGenOther.Columns.Info+" = ?", req.Info)
            }         

            if req.Img != "" {
                model = model.Where(dao.DemoGenOther.Columns.Img+" = ?", req.Img)
            }         

            if req.Imgs != "" {
                model = model.Where(dao.DemoGenOther.Columns.Imgs+" = ?", req.Imgs)
            }         

            if req.File != "" {
                model = model.Where(dao.DemoGenOther.Columns.File+" = ?", req.File)
            }         

            if req.Files != "" {
                model = model.Where(dao.DemoGenOther.Columns.Files+" = ?", req.Files)
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
func (s *demoGenOther) GetInfoById(id int64,ctx context.Context) (info *model.DemoGenOther, err error) {
	if id == 0 {
		err = gerror.New("参数错误")
		return
	}
	err = dao.DemoGenOther.Ctx(ctx).Where(dao.DemoGenOther.Columns.Id, id).Scan(&info)
	if err != nil {
		g.Log().Error(err)
	}
	if info == nil || err != nil {
		err = gerror.New("获取信息失败")
	}
	return
}

// Add 添加
func (s *demoGenOther) Add(req *dao.DemoGenOtherAddReq,ctx context.Context) (err error) {
	_, err = dao.DemoGenOther.Ctx(ctx).Insert(req)
	return
}

// Edit 修改
func (s *demoGenOther) Edit(req *dao.DemoGenOtherEditReq,ctx context.Context) error {    

	_, err := dao.DemoGenOther.Ctx(ctx).FieldsEx(dao.DemoGenOther.Columns.Id).Where(dao.DemoGenOther.Columns.Id, req.Id).
		Update(req)
	return err
}

// DeleteByIds 删除
func (s *demoGenOther) DeleteByIds(ids []int,ctx context.Context) (err error) {
	if len(ids) == 0 {
		err = gerror.New("参数错误")
		return
	}	

	_, err = dao.DemoGenOther.Ctx(ctx).Delete(dao.DemoGenOther.Columns.Id+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除失败")
	}
	return
}

