package models

import (
	"github.com/astaxie/beego/orm"
)

type Zones struct {
	Id       int    `orm:"pk"`
	Zonename string `乡镇名称`
}

func (this *Zones) TableName() string {
	return "zones"
}

func init() {
	orm.RegisterModel(new(Zones))
}
