package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webproject/4050/common/hjwt"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"webproject/4050/common/function"

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
	Adminid     string `核查员`
	Adminname   string `核查员姓名`
	Isyears     string `年度审核`
	Quarter1    string `一季度`
	Quarter2    string `二季度`
	Quarter3    string `三季度`
	Quarter4    string `四季度`
	Postion     string `签到位置`
	Photos      string `签到照片`
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
	qb.Select("a.id,a.username,a.realname,a.phone,a.avatarurl,a.worktype,a.updatetime,a.remark,b.zonename,a.isverify").
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
	qbs.Select("a.id,a.username,a.realname,a.phone,a.avatarurl,a.worktype,a.updatetime,a.remark,b.zonename,a.isverify").
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

func (this *VerifyController) Setremark() {
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
	qb.Select("a.id,a.username,a.realname,a.sex,a.address,a.bothtime,a.phone,a.worktype,a.workaddress,a.addtime,a.updatetime,a.remark,b.zonename,isverify").
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

type ApplysList struct {
	Id          int    `orm:"pk"`
	Userid      int    `用户ID`
	Realname    string `用户姓名`
	Years       string `年度`
	Workaddress string `就业地址`
	Worktype    string `就业形式`
	Isverify    int    `是否审核用户`
	Remark      string `备注`
	Adminid     int    `核查员`
	Adminname   string `核查员姓名`
	Isyears     int    `年度审核`
	Quarter1    int    `一季度`
	Quarter2    int    `二季度`
	Quarter3    int    `三季度`
	Quarter4    int    `四季度`
	Postion     string `签到位置`
	Photos      string `签到照片`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (this *VerifyController) Applys() {
	userid := this.GetString("userid")
	where := "1=1"
	o := orm.NewOrm()
	var maps []ApplysList
	//fmt.Println(id)

	if userid != "" {
		where = where + " and a.userid = " + userid
	}

	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.remark,a.isverify,a.isyears,a.quarter1,a.quarter2,a.quarter3,a.quarter4,a.postion,a.photos,a.adminid,b.realname as adminname").
		From("applys as a").
		LeftJoin("members as b").
		On("a.adminid = b.id").
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
func (this *VerifyController) Applyview() {
	id, _ := strconv.Atoi(this.GetString("id"))
	uid, _ := strconv.Atoi(this.GetString("uid"))
	o := orm.NewOrm()
	applys := new(ApplysList)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.isverify,a.isyears,a.quarter1,a.quarter2,a.quarter3,a.quarter4,a.postion,a.photos,a.adminid,b.realname as adminname,c.realname").
		From("applys as a").
		LeftJoin("members as b").
		On("a.adminid = b.id").
		LeftJoin("members as c").
		On("a.userid = c.id").
		Where("a.id = ?")
	sql := qb.String()
	err := o.Raw(sql, id).QueryRow(&applys)
	if err != nil {
		beego.Error(err)
	}
	this.Data["applys"] = applys
	this.Data["uid"] = uid
	this.TplName = "admin/verify_applyview.html"
}

type TongjiList struct {
	Id          int    `orm:"pk"`
	Userid      int    `用户ID`
	Username    string `用户名`
	Realname    string `真实姓名`
	Phone       string `电话`
	Years       string `年度`
	Workaddress string `就业地址`
	Worktype    string `就业形式`
	Isverify    int    `是否审核用户`
	Remark      string `备注`
	Adminid     int    `核查员`
	Adminname   string `核查员姓名`
	Isyears     int    `年度审核`
	Quarter1    int    `一季度`
	Quarter2    int    `二季度`
	Quarter3    int    `三季度`
	Quarter4    int    `四季度`
	Postion     string `签到位置`
	Photos      string `签到照片`
	Addtime     int64  `添加时间`
	Updatetime  int64  `更新时间`
}

func (this *VerifyController) TongjiList() {
	cookie := this.Ctx.GetCookie("Authorization")
	Claims, _ := hjwt.CheckToken(cookie)
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	sort := this.GetString("sortColumn")
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	where := "d.role_id is null"
	o := orm.NewOrm()
	var maps []TongjiList
	//	fmt.Println(id)
	zone_id := Claims["zone"].(float64)
	zone := strconv.FormatFloat(zone_id, 'f', 0, 64)
	role_id := Claims["role_id"].(float64)
	fmt.Println(zone_id)
	if role_id == 2 {
		where = where + " and b.id = " + zone
	}
	if search != "" {
		where = where + " and (c.username like '%" + search + "%' or c.realname like '%" + search + "%' or c.phone like '%" + search + "%' or b.realname like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.remark,a.isverify,a.isyears,a.quarter1,a.quarter2,a.quarter3,a.quarter4,a.postion,a.photos,a.adminid,b.realname as adminname,c.username,c.realname,c.phone").
		From("applys as a").
		LeftJoin("members as b").
		On("a.adminid = b.id").
		LeftJoin("members as c").
		On("a.userid = c.id").
		LeftJoin("role_member as d").
		On("a.userid = d.user_id").
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
	var counts []TongjiList
	qbs.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.remark,a.isverify,a.isyears,a.quarter1,a.quarter2,a.quarter3,a.quarter4,a.postion,a.photos,a.adminid,b.realname as adminname,c.username,c.realname,c.phone").
		From("applys as a").
		LeftJoin("members as b").
		On("a.adminid = b.id").
		LeftJoin("members as c").
		On("a.userid = c.id").
		LeftJoin("role_member as d").
		On("a.userid = d.user_id").
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

func (this *VerifyController) Outport() {
	cookie := this.Ctx.GetCookie("Authorization")
	Claims, _ := hjwt.CheckToken(cookie)

	search := this.GetString("search")
	where := "d.role_id is null"
	o := orm.NewOrm()
	var maps []TongjiList
	//	fmt.Println(id)
	zone_id := Claims["zone"].(float64)
	zone := strconv.FormatFloat(zone_id, 'f', 0, 64)
	role_id := Claims["role_id"].(float64)
	fmt.Println(zone_id)
	if role_id == 2 {
		where = where + " and b.id = " + zone
	}
	if search != "" {
		where = where + " and (c.username like '%" + search + "%' or c.realname like '%" + search + "%' or c.phone like '%" + search + "%' or b.realname like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.userid,a.years,a.addtime,a.updatetime,a.worktype,a.workaddress,a.remark,a.isverify,a.isyears,a.quarter1,a.quarter2,a.quarter3,a.quarter4,a.postion,a.photos,a.adminid,b.realname as adminname,c.username,c.realname,c.phone").
		From("applys as a").
		LeftJoin("members as b").
		On("a.adminid = b.id").
		LeftJoin("members as c").
		On("a.userid = c.id").
		LeftJoin("role_member as d").
		On("a.userid = d.user_id").
		Where(where).
		OrderBy("a.updatetime").Desc()

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)

	file, err := xlsx.OpenFile("moban.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet := file.Sheets[0]
	for _, value := range maps {

		row := sheet.AddRow()
		row.SetHeightCM(1)
		cell := row.AddCell()
		cell.Value = strconv.Itoa(value.Id)
		cell = row.AddCell()
		cell.Value = value.Years
		cell = row.AddCell()
		cell.Value = value.Username
		cell = row.AddCell()
		cell.Value = value.Realname
		cell = row.AddCell()
		cell.Value = value.Phone
		cell = row.AddCell()
		cell.Value = value.Worktype
		cell = row.AddCell()
		cell.Value = value.Workaddress
		cell = row.AddCell()
		if value.Isverify > 0 {
			cell.Value = "通过"
		} else if value.Isverify == -1 {
			cell.Value = "驳回"
		} else {
			cell.Value = "待审"
		}
		cell = row.AddCell()
		if value.Isyears > 0 {
			cell.Value = "通过"
		} else if value.Isyears == -1 {
			cell.Value = "不通过"
		} else {
			cell.Value = "待审"
		}
		cell = row.AddCell()
		if value.Quarter2 > 0 {
			cell.Value = "通过"
		} else if value.Quarter2 == -1 {
			cell.Value = "不通过"
		} else {
			cell.Value = "待审"
		}
		cell = row.AddCell()
		if value.Quarter3 > 0 {
			cell.Value = "通过"
		} else if value.Quarter3 == -1 {
			cell.Value = "不通过"
		} else {
			cell.Value = "待审"
		}
		cell = row.AddCell()
		if value.Quarter4 > 0 {
			cell.Value = "通过"
		} else if value.Quarter4 == -1 {
			cell.Value = "不通过"
		} else {
			cell.Value = "待审"
		}
		cell = row.AddCell()
		cell.Value = value.Postion
		cell = row.AddCell()
		cell.Value = value.Remark
		cell = row.AddCell()
		cell.Value = value.Adminname
		cell = row.AddCell()
		cell.Value = function.ConvertT(value.Updatetime)
		cell = row.AddCell()
	}
	err = file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}
	this.Ctx.Output.Header("Accept-Ranges", "bytes")
	this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s", "file.xlsx")) //文件名
	this.Ctx.Output.Header("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	this.Ctx.Output.Header("Pragma", "no-cache")
	this.Ctx.Output.Header("Expires", "0")
	//最主要的一句
	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, "file.xlsx")

}
