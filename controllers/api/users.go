package api

import (
	"fmt"
	"strconv"
	"webproject/4050/controllers"
	"webproject/4050/models"

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

func (this *UsersController) UserList() {
	id := this.GetString("id")
	limit := "10"
	istart := 0
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	where := "b.role_id IS NULL"
	o := orm.NewOrm()
	var maps []models.Members
	fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (username like '%" + search + "%' or realname like '%" + search + "%' or phone like '%" + search + "%' or email like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.username,a.realname,a.email,a.phone").
		From("members as a").
		LeftJoin("role_member as b").
		On("a.id=b.user_id").
		Where(where).
		OrderBy("a.id").Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	data := map[string]interface{}{"code": 1, "message": "success", "data": maps, "num": num}
	this.Data["json"] = data
	this.ServeJSON()
}
