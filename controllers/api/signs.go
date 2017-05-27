package api

import (
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
	signs := new(models.Signs)
	//fmt.Println(password)
	userid, _ := strconv.Atoi(this.GetString("userid"))
	isverify, _ := strconv.Atoi(this.GetString("isverify"))
	signs.Years = this.GetString("years")
	signs.Months = this.GetString("months")
	signs.Userid = userid
	signs.Postion = this.GetString("postion")
	signs.Photos = this.GetString("photos")
	signs.Isverify = isverify
	signs.Addtime = time.Now().Unix()
	signs.Updatetime = time.Now().Unix()
	id, err := o.Insert(signs)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "0", "message": "fail!"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "success!", "data": id}
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
