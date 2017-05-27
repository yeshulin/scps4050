package api

import (
	"fmt"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego/orm"
)

type ZonesController struct {
	controllers.ApibaseController
}

func (this *ZonesController) Get() {
	o := orm.NewOrm()
	var maps []models.Zones
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("*").
		From("zones").
		OrderBy("id").Asc()

	// 导出 SQL 语句
	sql := qb.String()

	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	data := map[string]interface{}{"data": maps, "code": "1", "message": "success!"}

	this.Data["json"] = data
	this.ServeJSON()
}
