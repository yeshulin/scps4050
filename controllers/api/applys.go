package api

import (
	"fmt"

	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego/orm"
)

type ApplysController struct {
	controllers.ApibaseController
}

func (this *ApplysController) List() {
	userid := this.GetString("userid")
	where := "1=1"
	o := orm.NewOrm()
	var maps []models.Applys
	//fmt.Println(id)

	if userid != "" {
		where = where + " and a.userid = " + userid
	}

	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.isverify").
		From("applys as a").
		Where(where).
		OrderBy("a.id").Desc()

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)
	num, _ := o.Raw(sql).QueryRows(&maps)

	data := map[string]interface{}{"code": 0, "message": "success", "data": maps, "total": num}
	this.Data["json"] = data
	this.ServeJSON()
}
