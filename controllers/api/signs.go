package api

import (
	"fmt"
	"strconv"
	"time"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego/orm"
)

type SignsController struct {
	controllers.ApibaseController
}

func (this *SignsController) Post() {
	o := orm.NewOrm()
	userid, _ := strconv.Atoi(this.GetString("userid"))
	isverify, _ := strconv.Atoi(this.GetString("isverify"))
	years := this.GetString("years")
	months := this.GetString("months")
	signsinfo := models.Signs{Userid: userid, Years: years, Months: months}
	err := o.Read(&signsinfo, "userid", "years", "months")
	if err != nil {
		signs := new(models.Signs)
		//fmt.Println(password)
		signs.Years = years
		signs.Months = months
		signs.Userid = userid
		signs.Postion = this.GetString("longitude") + "," + this.GetString("latitude")
		signs.Photos = this.GetString("photos")
		signs.Isverify = isverify
		signs.Addtime = time.Now().Unix()
		signs.Updatetime = time.Now().Unix()
		id, err := o.Insert(signs)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "提交申请失败!"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "你已经提交过本月底签到!"}
	}
	this.ServeJSON()
}

func (this *SignsController) Get() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	signs := models.Signs{Id: id}
	err := o.Read(&signs)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "签到记录找到!"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": signs}
	}
	this.ServeJSON()
}

func (this *SignsController) Put() {
	id, _ := strconv.Atoi(this.GetString("id"))
	isverify, _ := strconv.Atoi(this.GetString("isverify"))
	o := orm.NewOrm()
	signs := models.Signs{Id: id}
	err := o.Read(&signs)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "签到记录未找到!"}
	} else {
		signs.Isverify = isverify
		signs.Updatetime = time.Now().Unix()
		num, err1 := o.Update(&signs)
		if err1 != nil {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "审核失败"}
		} else {
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
		}
	}
	this.ServeJSON()
}

type SignsList struct {
	Id         int    `orm:"pk"`
	Years      int    `年度`
	Months     string `月度`
	Userid     int    `用户ID`
	Realname   string `真实姓名`
	Zonename   string `区域名称`
	Postion    string `位置`
	Photos     string `图片`
	Isverify   int    `是否审核`
	Remark     string `备注`
	Addtime    int64  `添加时间`
	Updatetime int64  `更新时间`
}

func (this *SignsController) List() {
	id := this.GetString("id")
	years := this.GetString("years")
	userid := this.GetString("userid")
	where := "1=1"
	o := orm.NewOrm()
	var maps []SignsList
	//fmt.Println(id)
	if id != "" {
		where = where + " and a.id = " + id
	}
	if years != "" {
		where = where + " and a.years = " + years
	}
	if userid != "" {
		where = where + " and a.userid = " + userid
	}

	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.years,a.userid,a.months,a.addtime,a.updatetime,a.photos,a.postion,a.isverify,b.realname,c.zonename").
		From("signs as a").
		LeftJoin("members as b").
		On("a.userid = b.id").
		LeftJoin("zones as c").
		On("b.zone = c.id").
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
