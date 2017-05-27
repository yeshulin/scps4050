package models

import (
	"github.com/astaxie/beego/orm"
)

type Members struct {
	Id          int    `orm:"pk"`
	Openid      string `openid`
	Username    string `用户名`
	Password    string `密码`
	Realname    string `真实姓名`
	Avatarurl   string `用户头像`
	Sex         string `性别`
	Bothtime    string `出生时间`
	Zone        int    `区域`
	Address     string `地址`
	Email       string `邮箱`
	Workaddress string `就业地址`
	Worktype    string `就业形式`
	Phone       string `电话`
	Isverify    int    `是否审核用户`
	Remark      string `备注`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (this *Members) TableName() string {
	return "members"
}

func init() {
	//orm.RegisterModel(new(Users), new(UsersProfile))
	orm.RegisterModel(new(Members))

}
