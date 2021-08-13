// ==========================================================================
// GFast自动生成model代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-08-13 15:17:57
// 生成路径: gfast/app/system/model/demo_gen_other.go
// 生成人：gfast
// ==========================================================================

package model

// DemoGenOther is the golang structure for table demo_gen_other.
type DemoGenOther struct {	

         Id       uint         `orm:"id,primary" json:"id"`    // ID    

         Info    string         `orm:"info" json:"info"`    // 内容    

         Img    string         `orm:"img" json:"img"`    // 单图    

         Imgs    string         `orm:"imgs" json:"imgs"`    // 多图    

         File    string         `orm:"file" json:"file"`    // 单文件    

         Files    string         `orm:"files" json:"files"`    // 多文件    

         Remark    string         `orm:"remark" json:"remark"`    // 描述    

}