// ==========================================================================
// GFast自动生成dao操作代码，无需手动修改，重新生成不会自动覆盖.
// 生成日期：2021-08-13 15:17:57
// 生成路径: gfast/app/system/dao/demo_gen_other.go
// 生成人：gfast
// ==========================================================================

package dao

import (
    comModel "gfast/app/common/model"
    "gfast/app/system/dao/internal"    

)

// demoGenOtherDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type demoGenOtherDao struct {
	*internal.DemoGenOtherDao
}

var (
    // DemoGenOther is globally public accessible object for table tools_gen_table operations.
    DemoGenOther = demoGenOtherDao{
        internal.NewDemoGenOtherDao(),
    }
)

// Fill with you ideas below.

// DemoGenOtherSearchReq 分页请求参数
type DemoGenOtherSearchReq struct {    

    Info  string `p:"info"` //内容    

    Img  string `p:"img"` //单图    

    Imgs  string `p:"imgs"` //多图    

    File  string `p:"file"` //单文件    

    Files  string `p:"files"` //多文件    

    comModel.PageReq
}

// DemoGenOtherAddReq 添加操作请求参数
type DemoGenOtherAddReq struct {    

    Info  string   `p:"info" `    

    Img  string   `p:"img" `    

    Imgs  string   `p:"imgs" `    

    File  string   `p:"file" `    

    Files  string   `p:"files" `    

    Remark  string   `p:"remark" `    

}

// DemoGenOtherEditReq 修改操作请求参数
type DemoGenOtherEditReq struct {
    Id    uint  `p:"id" v:"required#主键ID不能为空"`    

    Info  string `p:"info" `    

    Img  string `p:"img" `    

    Imgs  string `p:"imgs" `    

    File  string `p:"file" `    

    Files  string `p:"files" `    

    Remark  string `p:"remark" `    

}

