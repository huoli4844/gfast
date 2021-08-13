/*
* @desc:系统参数设置
* @company:云南省奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2021/7/5 18:00
 */

package service

import (
	"fmt"
	"gfast/app/common/dao"
	"gfast/app/common/global"
	"gfast/app/common/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type sysConfig struct {
}

var SysConfig = new(sysConfig)

func (s *sysConfig) SelectListByPage(req *model.SysConfigSearchReq) (total, page int, list []*model.SysConfig, err error) {
	m := dao.SysConfig.Ctx(req.Ctx)
	if req != nil {
		if req.ConfigName != "" {
			m = m.Where("config_name like ?", "%"+req.ConfigName+"%")
		}
		if req.ConfigType != "" {
			m = m.Where("config_type = ", gconv.Int(req.ConfigType))
		}
		if req.ConfigKey != "" {
			m = m.Where("config_key like ?", "%"+req.ConfigKey+"%")
		}
		if req.BeginTime != "" {
			m = m.Where("create_time >= ? ", req.BeginTime)
		}

		if req.EndTime != "" {
			m = m.Where("create_time<=?", req.EndTime)
		}
		if req.AppId != "" {
			m = m.Where("app_id=?", gconv.Int(req.AppId))
		}
	}
	total, err = m.Count()
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
		req.PageSize = model.PageSize
	}
	err = m.Page(page, req.PageSize).Order("config_id asc").Scan(&list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	return
}

// CheckConfigKeyUniqueAll 验证参数键名是否存在
func (s *sysConfig) CheckConfigKeyUniqueAll(configKey string) error {
	entity, err := dao.SysConfig.Fields(dao.SysConfig.C.ConfigId).FindOne(dao.SysConfig.C.ConfigKey, configKey)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("校验数据失败")
	}
	if entity != nil {
		return gerror.New("参数键名已经存在")
	}
	return nil
}

// AddSave 添加操作
func (s *sysConfig) AddSave(req *model.SysConfigAddReq) (err error) {
	_, err = dao.SysConfig.Insert(req)
	return
}

func (s *sysConfig) GetById(id int) (data *model.SysConfig, err error) {
	err = dao.SysConfig.WherePri(id).Scan(&data)
	return
}

// CheckConfigKeyUnique 检查键是否已经存在
func (s *sysConfig) CheckConfigKeyUnique(configKey string, configId int64) error {
	entity, err := dao.SysConfig.Fields(dao.SysConfig.C.ConfigId).
		FindOne(dao.SysConfig.C.ConfigKey+"=? and "+dao.SysConfig.C.ConfigId+"!=?",
			configKey, configId)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("校验数据失败")
	}
	if entity != nil {
		return gerror.New("参数键名已经存在")
	}
	return nil
}

// EditSave 修改系统参数
func (s *sysConfig) EditSave(req *model.SysConfigEditReq) (err error) {
	_, err = dao.SysConfig.FieldsEx(dao.SysConfig.C.ConfigId, dao.SysConfig.C.CreateBy).
		WherePri(req.ConfigId).Data(req).Update()
	return
}

// DeleteByIds 删除
func (s *sysConfig) DeleteByIds(ids []int) error {
	_, err := dao.SysConfig.Delete(dao.SysConfig.C.ConfigId+" in (?)", ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
	}
	return nil
}

//TODO：由于系统现在支持多公司都有一份自己的参数，所以登陆时候，需要通过登陆者归属的公司来获取它公司的对应的参数
//处理思路： 1、用户登陆时候，获取用户的部门信息 ，然后找到这个部门的顶级部门，即找到他的【A集团】这样的节点的部门，因为参数是以这个部门去保存到sys_org_config表中的
//相应的，sys_org_config表内容改以后，就删除缓存中的信息，让系统重新生成
// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sysConfig) GetConfigByKey(key string) (config *model.SysConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cache := Cache.New()
	cf := cache.Get(global.SysConfigTag + key)
	if cf != nil {
		err = gconv.Struct(cf, &config)
		return
	}
	config, err = s.GetByKey(key)
	if err != nil {
		return
	}
	cache.Set(global.SysConfigTag+key, config, 0, global.SysConfigTag)
	return
}

//TODO：由于系统现在支持多公司都有一份自己的参数，所以登陆时候，需要通过登陆者归属的公司来获取它公司的对应的参数

// GetByKey 通过key获取参数（从数据库获取）
func (s *sysConfig) GetByKey(key string) (config *model.SysConfig, err error) {
	err = dao.SysConfig.Where("config_key", key).Scan(&config)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取配置失败")
	}
	return
}

type dept struct {
	DeptId   int
	DeptName string
}

// SyncDataToOrgConfig  同步所有系统参数到 公司参数表中
/*
  处理步骤
  	获取所有的公司节点放到 deptLists 中
    获取所有的 sysconfig 到 paramList 中
  	获取所有的公司配置信息到 orgParamList 中

	一、循环 deptlists
		1.1 看看这个deptLists.dept_id 在 sys_org_config 中是否都存在，哪个公司不存在的，就用 insert into sys_org_config  (XXX,xxx) select (XXX,xxx) from sys_config 把数据完整的插入
		1.2 如果sys_org_config 中存在了这个公司的参数，那么扫描这个公司的参数是不是齐全 ，没有的在插入
*/
func (s *sysConfig) SyncDataToOrgConfig() error {
	//s := "SELECT * FROM `user` WHERE `status` IN(?)"
	//m := g.DB().Raw(s, g.Slice{1,2,3}).WhereLT("age", 18).Limit(10).OrderAsc("id").All()
	// SELECT * FROM `user` WHERE `status` IN(1,2,3) AND `age`=18 ORDER BY `id` ASC LIMIT 10

	//只检索出上级是 平台根目录的节点，这种节点就是公司的节点
	var deptLists []*dept
	sql := "select dept_id,dept_name from `sys_dept` WHERE `status` IN(?) and `parent_id`=100 "
	err := g.DB().Raw(sql, g.Slice{1}).OrderAsc("dept_id").Scan(&deptLists)

	if err != nil {
		g.Log().Error(err)
	}

	var paramList []*model.SysConfig
	sql = "select * from `sys_config`"
	err = g.DB().Raw(sql).OrderAsc("config_id").Scan(&paramList)
	if err != nil {
		g.Log().Error(err)
	}

	for _, dept := range deptLists {
		sql := ` select dept_id from sys_org_config where dept_id=%d`
		sql = fmt.Sprintf(sql, dept.DeptId)
		d1, _ := g.DB().Raw(sql).All()
		if len(d1) > 0 {
			//代表这个公司的数据已经存在， 那么就判断这个公司的参数是否齐全
			for _, param := range paramList {
				sql1 := ` select dept_id from sys_org_config where (dept_id=%d) and (config_id = %d)`
				sql1 = fmt.Sprintf(sql1, dept.DeptId, param.ConfigId)
				d2, _ := g.DB().Raw(sql1).All()
				if len(d2) > 0 {
					//代表这个公司的这个参数数据已经存在 那么就不处理
					continue
				} else {
					//把这个公司缺失的这个参数插入到这个公司中
					insertDataSql := `INSERT INTO sys_org_config (dept_id,config_id, config_name,config_key,config_value,config_type,remark,app_id) 
                      SELECT %d, config_id, config_name,config_key,config_value,config_type,remark,app_id FROM sys_config where config_id=%d`
					insertDataSql = fmt.Sprintf(insertDataSql, dept.DeptId, param.ConfigId)
					_, err = g.DB().Exec(insertDataSql)
					if err != nil {
						g.Log().Error(err)
					}

				}

			}
		} else {
			//这个公司没有数据，那么就插入一份完整的系统参数给这个公司
			insertSql := `INSERT INTO sys_org_config (dept_id,config_id, config_name,config_key,config_value,config_type,remark,app_id) 
                      SELECT %d, config_id, config_name,config_key,config_value,config_type,remark,app_id FROM sys_config`
			insertSql = fmt.Sprintf(insertSql, dept.DeptId)
			_, err = g.DB().Exec(insertSql)
			if err != nil {
				g.Log().Error(err)
			}
		}
	}

	return err
}
