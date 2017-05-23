package admin

import (
	"fmt"
	"strconv"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

type VerifyController struct {
	controllers.WebController
}

func (this *VerifyController) VerifyList() {
	this.TplName = "admin/verify_list.html"
}

type VerifyUser struct {
	Id          int    `orm:"pk"`
	Username    string `用户名`
	Realname    string `真实姓名`
	Sex         string `性别`
	Bothtime    string `出生时间`
	Zone        int    `区域`
	Zonename    string `区域名称`
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

func (this *VerifyController) Get() {
	id := this.GetString("id")
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	sort := this.GetString("sortColumn")
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	where := "c.role_id is null"
	o := orm.NewOrm()
	var maps []VerifyUser
	fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (username like '%" + search + "%' or realname like '%" + search + "%' or phone like '%" + search + "%' or email like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.username,a.realname,a.phone,a.worktype,a.updatetime,b.zonename,isverify").
		From("members as a").
		LeftJoin("zones as b").
		On("a.zone = b.id").
		LeftJoin("role_member as c").
		On("a.id = c.user_id").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []VerifyUser
	qbs.Select("a.id,a.username,a.realname,a.phone,a.worktype,a.updatetime,b.zonename,isverify").
		From("members as a").
		LeftJoin("zones as b").
		On("a.zone = b.id").
		LeftJoin("role_member as c").
		On("a.id = c.user_id").
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
func (this *VerifyController) Pass() {

	o := orm.NewOrm()
	member := new(models.Members)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		member.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(member) == nil {
			member.Isverify = 1
			id, err := o.Update(member)
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

func (this *VerifyController) Reject() {

	o := orm.NewOrm()
	member := new(models.Members)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		member.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(member) == nil {
			member.Isverify = -1
			member.Remark = this.GetString("remark")
			id, err := o.Update(member)
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
