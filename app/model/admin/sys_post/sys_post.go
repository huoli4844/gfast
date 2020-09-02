// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package sys_post

import (
	"database/sql"
	"gfast/library/service"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

// Fill with you ideas below.

type Post struct {
	Entity
	SearchValue interface{} `json:"searchValue"`
	Remark      string      `json:"remark"`
	DataScope   interface{} `json:"dataScope"`
	Params      struct{}    `json:"params"`
	Flag        bool        `json:"flag"`
}

type SearchParams struct {
	PageNum  int    `p:"page"`     //当前页码
	PageSize int    `p:"pageSize"` //每页数
	PostCode string `p:"postCode"` //岗位编码
	PostName string `p:"postName"` //岗位名称
	Status   string `p:"status"`   //状态
}

type AddParams struct {
	PostCode string `p:"postCode" v:"required#岗位编码不能为空"`
	PostName string `p:"postName" v:"required#岗位名称不能为空"`
	PostSort int    `p:"postSort" v:"required#岗位排序不能为空"`
	Status   string `p:"status" v:"required#状态不能为空"`
	Remark   string `p:"remark"`
	AddUser  int
}

type EditParams struct {
	PostId int64 `p:"postId" v:"required#id必须"`
	AddParams
	UpUser int
}

func List(req *SearchParams) (total, page int, list gdb.Result, err error) {
	model := g.DB().Table(Table)

	if req != nil {
		if req.PostCode != "" {
			model.Where("post_code like ?", "%"+req.PostCode+"%")
		}

		if req.PostName != "" {
			model.Where("post_name like ?", "%"+req.PostName+"%")
		}

		if req.Status != "" {
			model.Where("status", req.Status)
		}
	}

	total, err = model.Count()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
	}

	if req.PageNum == 0 {
		req.PageNum = 1
	}

	page = req.PageNum

	if req.PageSize == 0 {
		req.PageSize = service.AdminPageNum
	}

	list, err = model.Page(page, req.PageSize).All()

	//g.Log().Println(list)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}

	return
}

//获取正常状态的岗位
func GetUsedPost() (list []*Entity, err error) {
	list, err = Model.Where(Columns.Status, 1).Order(Columns.PostSort + " ASC, " + Columns.PostId + " ASC ").All()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取岗位数据失败")
	}
	return
}

/**
添加
*/
func Add(addParams *AddParams) (result sql.Result, err error) {
	entity := &Entity{
		PostCode:   addParams.PostCode,
		PostName:   addParams.PostName,
		PostSort:   addParams.PostSort,
		Status:     addParams.Status,
		Remark:     addParams.Remark,
		CreateBy:   gconv.String(addParams.AddUser),
		CreateTime: gtime.Now(),
	}

	return entity.Save()

}

func Edit(editParams *EditParams) (result sql.Result, err error) {
	var entity *Entity
	entity, err = Model.FindOne(editParams.PostId)
	entity.PostId = editParams.PostId
	entity.PostCode = editParams.PostCode
	entity.PostName = editParams.PostName
	entity.PostSort = editParams.PostSort
	entity.Status = editParams.Status
	entity.Remark = editParams.Remark
	entity.UpdateBy = gconv.String(editParams.UpUser)
	entity.UpdateTime = gtime.Now()
	if entity.CreateTime == nil {
		entity.CreateTime = entity.UpdateTime
	}
	return entity.Update()
}

func GetOneById(id int64) (*Entity, error) {
	return Model.One(g.Map{
		"post_id": id,
	})
}

func DeleteByIds(ids []int) error {
	_, err := Model.Delete("post_id IN(?)", ids)
	if err != nil {
		g.Log().Error(err)
		return gerror.New("删除失败")
	}
	return nil
}
