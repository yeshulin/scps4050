package admin

import (
	"html/template"
	//"net/http"
	//"crypto/aes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
	"webproject/4050/common/aesencrypt"
	"webproject/4050/common/hjwt"
	"webproject/4050/controllers"
	"webproject/4050/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	//	jwt "github.com/dgrijalva/jwt-go"
	//	"github.com/astaxie/beego/validation"
)

type UserController struct {
	controllers.WebController
}

func (this *UserController) Login() {

	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/login.html"
}

func (this *UserController) Post() {
	aesencrypt := new(aesencrypt.AesEncrypt)
	username := this.GetString("account")
	o := orm.NewOrm()
	member := models.Members{Username: username}
	err := o.Read(&member, "username")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户不存在!"}
	}
	pass, _ := aesencrypt.Encrypt(this.GetString("password"))
	//	fmt.Println(base64.StdEncoding.EncodeToString(pass))
	//	fmt.Println(member.Password)

	if base64.StdEncoding.EncodeToString(pass) != member.Password {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户名密码错误!"}

	} else {
		var roleuser RoleUser
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("a.id,a.role_id,b.name as role_name,a.user_id,b.name as rolename,c.username,c.realname").
			From("role_member as a").
			LeftJoin("role as b").On("a.role_id = b.id").
			LeftJoin("members as c").On("a.user_id = c.id").
			Where("c.id = ?").
			OrderBy("a.id").Desc()

		// 导出 SQL 语句
		sql := qb.String()

		// 执行 SQL 语句

		o.Raw(sql, member.Id).QueryRow(&roleuser)
		fmt.Println(roleuser)
		if roleuser.Role_id < 1 {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户没有权限!"}
		} else {
			token := hjwt.GenToken(member.Id, member.Username, member.Realname, member.Email, member.Phone, member.Zone, roleuser.Role_id, roleuser.Rolename)
			this.Ctx.SetCookie("Authorization", token, 86400, "/")
			//		this.Ctx.Redirect(302, "/admin")
			this.Data["json"] = map[string]interface{}{"code": "1", "message": "success", "data": member.Username}
		}

	}
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//		"sub":      member.Id,
	//		"iat":      iat,
	//		"exp":      exp,
	//		"username": member.Username,
	//		"email":    member.Email,
	//		"phone":    member.Phone,
	//		"realname": member.Realname,
	//	})

	//	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("jwtkey")))
	//	//更新登录时间，用于只允许用户一台设备登录
	//	/*this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "tokenString": tokenString}
	//	this.ServeJSON()*/
	//	fmt.Println(tokenString)
	this.ServeJSON()

}
func (this *UserController) Reg() {
	aesencrypt := new(aesencrypt.AesEncrypt)
	password, err := aesencrypt.Encrypt(this.GetString("password"))
	if err != nil {
		beego.Error(err)
	}
	o := orm.NewOrm()
	member := new(models.Members)
	//fmt.Println(password)
	member.Username = this.GetString("username")
	member.Password = base64.StdEncoding.EncodeToString(password)
	member.Realname = this.GetString("realname")
	member.Email = this.GetString("email")
	member.Phone = this.GetString("phone")
	member.Addtime = time.Now().Unix()
	member.Updatetime = time.Now().Unix()
	id, err := o.Insert(member)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	}
	this.ServeJSON()
}

func (this *UserController) LoginWeixin() {
	aesencrypt := new(aesencrypt.AesEncrypt)
	username := this.GetString("username")
	o := orm.NewOrm()
	member := models.Members{Username: username}
	err := o.Read(&member, "username")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户不存在!"}
	}
	pass, _ := aesencrypt.Encrypt(this.GetString("password"))

	fmt.Println(base64.StdEncoding.EncodeToString(pass))
	fmt.Println(member.Password)
	if base64.StdEncoding.EncodeToString(pass) != member.Password {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户名密码错误!"}

	} else {
		var roleuser RoleUser
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("a.id,a.role_id,b.name as role_name,a.user_id,b.name as rolename,c.username,c.realname").
			From("role_member as a").
			LeftJoin("role as b").On("a.role_id = b.id").
			LeftJoin("members as c").On("a.user_id = c.id").
			Where("c.id = ?").
			OrderBy("a.id").Desc()

		// 导出 SQL 语句
		sql := qb.String()

		// 执行 SQL 语句

		o.Raw(sql, member.Id).QueryRow(&roleuser)
		fmt.Println(roleuser)
		if roleuser.Role_id < 1 {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户没有权限!"}
		} else {

			//this.Ctx.Redirect(302, "/admin")
			member1 := models.Members{Id: member.Id}
			if o.Read(&member1) == nil {
				member1.Openid = this.GetString("openid")
				member1.Avatarurl = this.GetString("avatarurl")
				if num, err := o.Update(&member1); err == nil {
					fmt.Println(num)
					this.Data["json"] = map[string]interface{}{"code": "1", "message": "success", "data": roleuser.Role_id}
				} else {
					this.Data["json"] = map[string]interface{}{"code": "0", "message": "用户登录失败!"}
				}
			}
		}

	}

	this.ServeJSON()

}

type WeixinResult struct {
	session_key string `会话密钥`
	openid      string `用户唯一标识`
	expires_in  int    `过期时间`
}

func (this *UserController) RegWeixin() {
	code := this.GetString("code")
	fmt.Println(code)
	AppID := beego.AppConfig.String("wxAppID")
	AppSecret := beego.AppConfig.String("wxAppSecret")
	httpUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + AppID + "&secret=" + AppSecret + "&js_code=" + code + "&grant_type=authorization_code"
	fmt.Println(httpUrl)
	result := make(map[string]interface{})
	req := httplib.Get(httpUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.ToJSON(&result)
	//	fmt.Println(result["session_key"])
	//	aesencrypt := new(aesencrypt.AesEncrypt)
	//	password, err := aesencrypt.Encrypt("123456")
	//	if err != nil {
	//		beego.Error(err)
	//	}
	o := orm.NewOrm()
	meminfo := models.Members{Openid: result["openid"].(string)}
	err := o.Read(&meminfo, "openid")
	if err != nil {
		//		member := new(models.Members)
		//		//fmt.Println(password)
		//		member.Openid = result["openid"].(string)
		//		member.Password = base64.StdEncoding.EncodeToString(password)
		//		member.Addtime = time.Now().Unix()
		//		member.Updatetime = time.Now().Unix()
		//		id, _ := o.Insert(member)
		//		result["id"] = id
		result["id"] = 0
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "success!", "data": result}
	} else {
		result["id"] = meminfo.Id
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "success!", "data": result}
	}
	this.ServeJSON()
}

func (this *UserController) Apply() {
	o := orm.NewOrm()

	zone, _ := strconv.Atoi(this.GetString("zone"))
	years := this.GetString("years")

	member1 := models.Members{Username: this.GetString("username"), Openid: ""}
	err := o.Read(&member1, "username")

	if err == orm.ErrNoRows {
		member := new(models.Members)
		aesencrypt := new(aesencrypt.AesEncrypt)
		password, _ := aesencrypt.Encrypt("123456")
		member.Openid = this.GetString("openid")
		member.Password = base64.StdEncoding.EncodeToString(password)
		member.Username = this.GetString("username")
		member.Realname = this.GetString("realname")
		member.Avatarurl = this.GetString("avatarurl")
		member.Sex = this.GetString("sex")
		member.Bothtime = this.GetString("bothtime")
		member.Zone = zone + 1
		member.Address = this.GetString("address")
		member.Email = this.GetString("email")
		member.Phone = this.GetString("phone")
		member.Workaddress = this.GetString("workaddress")
		member.Worktype = this.GetString("worktype")
		member.Addtime = time.Now().Unix()
		member.Updatetime = time.Now().Unix()
		if id, err := o.Insert(member); err == nil {
			applys := models.Applys{Userid: int(id), Years: years}
			err = o.Read(&applys, "userid", "years")
			if err != nil {
				applys1 := new(models.Applys)
				applys1.Userid = int(id)
				applys1.Years = this.GetString("years")
				applys1.Workaddress = this.GetString("workaddress")
				applys1.Worktype = this.GetString("worktype")
				applys1.Addtime = time.Now().Unix()
				applys1.Updatetime = time.Now().Unix()
				if id, err := o.Insert(applys1); err == nil {
					fmt.Println(id)
					this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
				} else {
					this.Data["json"] = map[string]interface{}{"code": "0", "message": "提交申请失败!"}
				}
			} else {
				this.Data["json"] = map[string]interface{}{"code": "0", "message": "你已经提交过本年度资料!"}
			}

		} else {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "提交申请失败!"}
		}
	} else {
		member := models.Members{Id: member1.Id}
		if o.Read(&member) == nil {
			member.Openid = this.GetString("openid")
			member.Avatarurl = this.GetString("avatarurl")
			if num, err := o.Update(&member); err == nil {
				this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
			} else {
				this.Data["json"] = map[string]interface{}{"code": "0", "message": "提交申请失败!"}
			}
		}
	}
	this.ServeJSON()
}

func (this *UserController) GetApply() {
	o := orm.NewOrm()
	userid, _ := strconv.Atoi(this.GetString("userid"))
	years := this.GetString("years")
	applys := models.Applys{Userid: userid, Years: years, Isverify: 1}
	err := o.Read(&applys, "userid", "years", "isverify")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "当前资料未审核不能签到!"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "可以签到!"}
	}
	this.ServeJSON()
}
func (this *UserController) Verify() {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(this.GetString("id"))
	isverify, _ := strconv.Atoi(this.GetString("isverify"))
	member := models.Members{Id: id}

	err := o.Read(&member)

	if err == orm.ErrNoRows {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "找不到用户!"}
	} else if err == orm.ErrMissPK {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "找不到用户!"}
	} else {
		member.Isverify = isverify
		member.Remark = this.GetString("remark")
		if num, err := o.Update(&member); err == nil {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "success!", "data": num}
		} else {
			this.Data["json"] = map[string]interface{}{"code": "0", "message": "审核失败!"}
		}
	}
	this.ServeJSON()
}

func (this *UserController) Register() {
	this.TplName = "admin/register.html"
}

func (this *UserController) UserList() {
	this.TplName = "admin/userlist.html"
}

type Counts struct {
	total int `总量`
}

type MemberList struct {
	Id       int    `orm:"pk"`
	Username string `用户名`
	Password string `密码`
	Realname string `真实姓名`
	Zone     int    `区域`
	Zonename string `区域名称`
	Email    string `邮箱`
	Phone    string `电话`
	Name     string `角色名`
}

func (this *UserController) Get() {
	id := this.GetString("id")
	limit := "10"
	start := this.GetString("start")
	page := this.GetString("page")
	sort := this.GetString("sortColumn")
	search := this.GetString("search")
	ilimit, _ := strconv.Atoi(limit)
	istart, _ := strconv.Atoi(start)
	where := "b.role_id != ''"
	o := orm.NewOrm()
	var maps []MemberList
	fmt.Println(id)
	if id != "" {
		where = where + " and id = " + id
	}
	if search != "" {
		where = where + " and (username like '%" + search + "%' or realname like '%" + search + "%' or phone like '%" + search + "%' or email like '%" + search + "%')"
	}
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("a.id,a.username,a.realname,a.email,a.phone,a.zone,c.name,d.zonename").
		From("members as a").
		LeftJoin("role_member as b").
		On("a.id=b.user_id").
		LeftJoin("role as c").
		On("b.role_id = c.id").
		LeftJoin("zones as d").
		On("a.zone = d.id").
		Where(where).
		OrderBy(sort).Desc().
		Limit(ilimit).Offset(istart)

	// 导出 SQL 语句
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&maps)
	fmt.Println(num)
	/*查询总量*/
	qbs, _ := orm.NewQueryBuilder("mysql")
	var counts []MemberList
	qbs.Select("a.id,a.username,a.realname,a.email,a.phone,a.zone,c.name,d.zonename").
		From("members as a").
		LeftJoin("role_member as b").
		On("a.id=b.user_id").
		LeftJoin("role as c").
		On("b.role_id = c.id").
		LeftJoin("zones as d").
		On("a.zone = d.id").
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

func (this *UserController) Add() {
	/*用户角色*/
	o := orm.NewOrm()
	qbs, _ := orm.NewQueryBuilder("mysql")
	var roles []models.Role
	qbs.Select("id,name,status,remark,addtime,updatetime").
		From("role").
		OrderBy("id").Asc()
	sqls := qbs.String()
	nums, _ := o.Raw(sqls).QueryRows(&roles)
	fmt.Println(nums)
	/*所属乡镇*/
	qb, _ := orm.NewQueryBuilder("mysql")
	var zones []models.Zones
	qb.Select("id,zonename").
		From("zones").
		OrderBy("id").
		Asc()
	sql := qb.String()
	num, _ := o.Raw(sql).QueryRows(&zones)
	fmt.Println(num)
	this.Data["zones"] = zones
	this.Data["roles"] = roles
	this.TplName = "admin/user_add.html"
}

func (this *UserController) AddPost() {
	fmt.Println(this.GetString("username"))
	aesencrypt := new(aesencrypt.AesEncrypt)
	password, err := aesencrypt.Encrypt(this.GetString("password"))
	if err != nil {
		beego.Error(err)
	}
	zone, _ := strconv.Atoi(this.GetString("zone"))
	o := orm.NewOrm()
	member := new(models.Members)
	//fmt.Println(password)
	member.Username = this.GetString("username")
	member.Password = base64.StdEncoding.EncodeToString(password)
	member.Realname = this.GetString("realname")
	member.Email = this.GetString("email")
	member.Phone = this.GetString("phone")
	member.Zone = zone
	member.Addtime = time.Now().Unix()
	member.Updatetime = time.Now().Unix()
	fmt.Println(member)
	id, err := o.Insert(member)
	if err != nil {
		beego.Error(err)
	}
	role_ids, _ := strconv.Atoi(this.GetString("role_id"))
	fmt.Println(role_ids)
	res, err := o.Raw("delete from role_member where user_id = ? and role_id in(?)", id, role_ids).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	rolemember := new(models.RoleMember)
	rolemember.User_id = int(id)
	rolemember.Role_id = role_ids
	_, err1 := o.Insert(rolemember)
	if err1 != nil {
		beego.Error(err1)
	}
	/*defer func() {
		this.Redirect("/login", 302)
	}()*/
	this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
	this.ServeJSON()

}

func (this *UserController) Delete() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	if num, err := o.Delete(member); err == nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": num}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	}
	this.ServeJSON()
}

func (this *UserController) View() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	err := o.Read(member)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["member"] = member
	this.TplName = "admin/user_view.html"
}

func (this *UserController) Edit() {
	id, _ := strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	member := new(models.Members)
	member.Id = id
	err := o.Read(member)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	this.Data["member"] = member
	this.TplName = "admin/user_edit.html"
}

func (this *UserController) EditPost() {
	fmt.Println(this.GetString("username"))

	o := orm.NewOrm()
	member := new(models.Members)

	//fmt.Println(password)
	id, err := strconv.Atoi(this.GetString("id"))
	fmt.Println(id)

	if err == nil {
		member.Id = id
		fmt.Println(id)
		//fmt.Println(o.Read(&member))
		if o.Read(member) == nil {
			member.Username = this.GetString("username")
			password := this.GetString("password")
			if password != "" {
				aesencrypt := new(aesencrypt.AesEncrypt)
				passwords, err := aesencrypt.Encrypt(password)
				if err != nil {
					beego.Error(err)
				}
				member.Password = base64.StdEncoding.EncodeToString(passwords)
			}
			member.Realname = this.GetString("realname")
			member.Email = this.GetString("email")
			member.Phone = this.GetString("phone")
			member.Addtime = time.Now().Unix()
			member.Updatetime = time.Now().Unix()
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

func (this *UserController) LogOut() {
	this.Ctx.SetCookie("Authorization", "", 86400, "/")
	this.Ctx.Redirect(302, "/login")
}
