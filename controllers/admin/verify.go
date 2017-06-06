package admin

import (
	"fmt"
	"strconv"
	"time"
	"webproject/4050/common/hjwt"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
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
	Avatarurl   string `用户头像`
	Worktype    string `就业形式`
	Phone       string `电话`
	Isverify    int    `是否审核用户`
	Remark      string `备注`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (this *VerifyController) Get() {
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
	where := "c.role_id is null"
	o := orm.NewOrm()
	var maps []VerifyUser
	//	fmt.Println(id)
	zone_id := Claims["zone"].(float64)
	zone := strconv.FormatFloat(zone_id, 'f', 0, 64)
	role_id := Claims["role_id"].(float64)
	fmt.Println(zone_id)
	if role_id == 2 {
		where = where + " and b.id = " + zone
	}
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (username like '%" + search + "%' or realname like '%" + search + "%' or phone like '%" + search + "%' or email like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.username,a.realname,a.phone,a.avatarurl,a.worktype,a.updatetime,b.zonename,a.isverify").
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
	qbs.Select("a.id,a.username,a.realname,a.phone,a.avatarurl,a.worktype,a.updatetime,b.zonename,a.isverify").
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
func (this *VerifyController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	member := new(VerifyUser)
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.username,a.realname,a.sex,a.address,a.bothtime,a.phone,a.worktype,a.workaddress,a.addtime,a.updatetime,b.zonename,isverify").
		From("members as a").
		LeftJoin("zones as b").
		On("a.zone = b.id").
		LeftJoin("role_member as c").
		On("a.id = c.user_id").
		Where("a.id = ?")
	sql := qb.String()
	err := o.Raw(sql, id).QueryRow(&member)
	if err != nil {
		beego.Error(err)
	}
	this.Data["member"] = member
	this.TplName = "admin/verify_view.html"
}

func (this *VerifyController) Tongji() {
	this.TplName = "admin/verify_tongji.html"
}

func (this *VerifyController) Applys() {
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
func (this *VerifyController) Passapplys() {

	o := orm.NewOrm()
	applys := new(models.Applys)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		applys.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(applys) == nil {
			applys.Isverify = 1
			id, err := o.Update(applys)
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

func (this *VerifyController) Rejectapplys() {

	o := orm.NewOrm()
	applys := new(models.Applys)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		applys.Id = id
		//		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(applys) == nil {
			applys.Isverify = -1
			applys.Remark = this.GetString("remark")
			id, err := o.Update(applys)
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

func (this *VerifyController) Import() {
	this.TplName = "admin/verify_import.html"
}

func (this *VerifyController) Importpost() {
	o := orm.NewOrm()
	excelFileName := this.GetString("xlsx")
	fmt.Println(excelFileName)
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		beego.Error(error)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			//			fmt.Println(row.Cells[0].String)
			//			val, _ := row.Cells[0].String()
			realname, _ := row.Cells[1].String()
			sex, _ := row.Cells[2].String()
			bothtime, _ := row.Cells[3].String()
			zonename, _ := row.Cells[4].String()
			fmt.Println(zonename)
			zones := models.Zones{Zonename: zonename}
			err := o.Read(&zones, "zonename")
			if err != nil {
				beego.Error(err)
			}
			zoneid := zones.Id
			address, _ := row.Cells[5].String()
			username, _ := row.Cells[6].String()
			phone, _ := row.Cells[7].String()
			workaddress, _ := row.Cells[8].String()
			worktype, _ := row.Cells[9].String()
			member := new(models.Members)
			member.Username = username
			member.Password = "RE7+OOlq"
			member.Realname = realname
			member.Sex = sex
			member.Bothtime = bothtime
			member.Zone = zoneid
			member.Phone = phone
			member.Address = address
			member.Isverify = 1
			member.Addtime = time.Now().Unix()
			member.Updatetime = time.Now().Unix()
			id, _ := o.Insert(member)
			fmt.Println(id)
			applys := new(models.Applys)
			applys.Userid = int(id)
			applys.Years = "2017"
			applys.Worktype = worktype
			applys.Workaddress = workaddress
			applys.Isverify = 1
			applys.Addtime = time.Now().Unix()
			applys.Updatetime = time.Now().Unix()
			id1, _ := o.Insert(applys)
			fmt.Println(id1)
			//			fmt.Println(val)
			//			for _, cell := range row.Cells {
			//				val, _ := cell.String()
			//				fmt.Printf("%s\n", val)
			//			}
		}
	}
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!"}
	this.ServeJSON()
}
