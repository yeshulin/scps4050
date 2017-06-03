package api

import (
	"fmt"
	"webproject/4050/controllers"

	"github.com/astaxie/beego/orm"
)

type UsersController struct {
	controllers.ApibaseController
}

type MembersApply struct {
	Id          int    `orm:"pk"`
	Openid      string `openid`
	Username    string `用户名`
	Realname    string `真实姓名`
	Sex         string `性别`
	Bothtime    string `出生时间`
	Zone        int    `区域`
	Address     string `地址`
	Email       string `邮箱`
	Workaddress string `就业地址`
	Worktype    string `就业形式`
	Phone       string `电话`
}

func (this *UsersController) Get() {
	userid := this.GetString("userid")
	o := orm.NewOrm()
	member := new(MembersApply)
	where := "1=1"
	if userid != "" {
		where = where + " and a.id = " + userid
	}

	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.openid,a.username,a.realname,a.sex,a.bothtime,a.zone,a.address,a.email,a.phone,b.worktype,b.workaddress").
		From("members as a").
		LeftJoin("applys as b").
		On("a.id = b.userid").
		Where(where).
		OrderBy("a.id").Desc()
	sql := qb.String()
	fmt.Println(sql)
	num := o.Raw(sql).QueryRow(&member)

	data := map[string]interface{}{"code": 1, "message": "success", "data": member, "num": num}
	this.Data["json"] = data
	this.ServeJSON()
}
