package admin

import (
	"fmt"
	"strconv"
	//	"time"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"webproject/4050/common/hjwt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SignsController struct {
	controllers.WebController
}

func (this *SignsController) SignsList() {
	this.TplName = "admin/signs_list.html"
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

func (this *SignsController) Get() {
	cookie := this.Ctx.GetCookie("Authorization")
	Claims, _ := hjwt.CheckToken(cookie)
	id := this.GetString("id")
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	sort := this.GetString("sortColumn")
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	where := "1=1"
	o := orm.NewOrm()
	var maps []SignsList
	//fmt.Println(id)
	zone_id := Claims["zone"].(float64)
	zone := strconv.FormatFloat(zone_id, 'f', 0, 64)
	role_id := Claims["role_id"].(float64)
	fmt.Println(zone_id)
	if role_id == 2 {
		where = where + " and c.id = " + zone
	}
	if id != "" {
		where = where + " and a.id = " + id
	}
	if search != "" {
		where = where + " and (a. like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.years,a.months,a.addtime,a.updatetime,a.postion,a.isverify,b.realname,c.zonename").
		From("signs as a").
		LeftJoin("members as b").
		On("a.userid = b.id").
		LeftJoin("zones as c").
		On("b.zone = c.id").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []SignsList
	qbs.Select("a.id,a.years,a.months,a.addtime,a.updatetime,a.postion,a.isverify,b.realname,c.zonename").
		From("signs as a").
		LeftJoin("members as b").
		On("a.userid = b.id").
		LeftJoin("zones as c").
		On("b.zone = c.id").
		Where(where).
		OrderBy(sort).Desc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&counts)
	fmt.Println(nums)
	//fmt.Println(counts)
	//fmt.Println(maps)
	data := map[string]interface{}{"data": maps, "limit": limit, "page": page, "total": nums}
	//data["data"] = maps
	//json.data = maps
	//json.limit = limit
	//json.page = page
	//json.total = nums

	fmt.Println(limit)
	fmt.Println(page)
	fmt.Println(data)
	this.Data["json"] = data
	this.ServeJSON()

}

func (this *SignsController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	signs := new(SignsList)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("a.id,a.years,a.months,a.addtime,a.updatetime,a.postion,a.photos,a.isverify,b.realname,c.zonename").
		From("signs as a").
		LeftJoin("members as b").
		On("a.userid = b.id").
		LeftJoin("zones as c").
		On("b.zone = c.id").
		Where("a.id = ?")
	sql := qb.String()
	err := o.Raw(sql, id).QueryRow(&signs)
	if err != nil {
		beego.Error(err)
	}
	this.Data["signs"] = signs
	this.TplName = "admin/signs_view.html"
}

func (this *SignsController) Pass() {

	o := orm.NewOrm()
	signs := new(models.Signs)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		signs.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&signs))
		if o.Read(signs) == nil {
			signs.Isverify = 1
			id, err := o.Update(signs)
			if err != nil {
				beego.Error(err)
			}
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}

		} else {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
		}
	}
	this.ServeJSON()
}

func (this *SignsController) Reject() {

	o := orm.NewOrm()
	signs := new(models.Signs)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		signs.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&signs))
		if o.Read(signs) == nil {
			signs.Isverify = -1
			signs.Remark = this.GetString("remark")
			id, err := o.Update(signs)
			if err != nil {
				beego.Error(err)
			}
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}

		} else {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
		}
	}
	this.ServeJSON()
}
