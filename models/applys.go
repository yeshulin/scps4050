package models

import (
	"github.com/astaxie/beego/orm"
)

type Applys struct {
	Id          int    `orm:"pk"`
	Userid      int    `用户ID`
	Years       string `年度`
	Workaddress string `就业地址`
	Worktype    string `就业形式`
	Isverify    int    `是否审核用户`
	Remark      string `备注`
	Adminid     int    `核查员`
	Isyears     int    `年度审核`
	Quarter1    int    `一季度`
	Quarter2    int    `二季度`
	Quarter3    int    `三季度`
	Quarter4    int    `四季度`
	Postion     string `签到位置`
	Photos      string `签到照片`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (this *Applys) TableName() string {
	return "applys"
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Applys))

}
