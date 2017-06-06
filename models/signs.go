package models

import (
	"github.com/astaxie/beego/orm"
)

type Signs struct {
	Id         int    `orm:"pk"`
	Years      string `年度`
	Quarter    string `季度`
	Months     string `月度`
	Userid     int    `用户ID`
	Postion    string `签到位置`
	Photos     string `签到照片`
	Isverify   int    `是否审核`
	Remark     string `反馈信息`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (this *Signs) TableName() string {
	return "signs"
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Signs))

}
