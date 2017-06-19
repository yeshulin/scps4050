package api

import (
	"fmt"
	"strconv"
	"time"
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

func (this *ApplysController) Post() {
	o := orm.NewOrm()
	userid, _ := strconv.Atoi(this.GetString("userid"))
	years := this.GetString("years")
	adminid, _ := strconv.Atoi(this.GetString("adminid"))
	quarter1, _ := strconv.Atoi(this.GetString("quarter1"))
	quarter2, _ := strconv.Atoi(this.GetString("quarter2"))
	quarter3, _ := strconv.Atoi(this.GetString("quarter3"))
	quarter4, _ := strconv.Atoi(this.GetString("quarter4"))
	applysinfo := models.Applys{Userid: userid, Years: years}
	err := o.Read(&applysinfo, "userid", "years")
	fmt.Println(applysinfo)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "资料记录未找到!"}
	} else {
		applys := models.Applys{Id: applysinfo.Id}
		err := o.Read(&applys)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "资料记录未找到!"}
		} else {
			applys.Isverify = 1
			applys.Adminid = adminid
			applys.Isyears = 1
			applys.Quarter1 = quarter1
			applys.Quarter2 = quarter2
			applys.Quarter3 = quarter3
			applys.Quarter4 = quarter4
			applys.Postion = this.GetString("longitude") + "," + this.GetString("latitude")
			applys.Photos = this.GetString("photos")
			applys.Updatetime = time.Now().Unix()
			num, err1 := o.Update(&applys)
			if err1 != nil {
				this.Data["json"] = map[string]interface{}{"code": "0", "message": "资料核查失败！"}
			} else {
				this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
			}
		}

	}
	this.ServeJSON()
}
