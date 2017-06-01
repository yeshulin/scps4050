package controllers

import (
	"fmt"
	//	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) Upload() {
	f, h, err := this.GetFile("file")
	if err != nil {
		json := map[string]interface{}{"code": "0", "message": "fail!"}
		this.Data["json"] = json
		beego.Error(err)
	} else {
		defer f.Close()

		//	year := strconv.Itoa(time.Now().Year())
		//	month := time.Now().Month().String()
		//	day := strconv.Itoa(time.Now().Day())
		// 获取当前年月
		datePath := time.Now().Format("2006/01/01")
		// 设置保存目录
		dirPath := beego.AppConfig.String("UploadPath") + datePath
		fmt.Println(h.Filename)

		this.SaveToFile("file", fmt.Sprintf("%s/%s", dirPath, h.Filename))
		json := map[string]interface{}{"code": "1", "message": "success!", "data": fmt.Sprintf("%s/%s", dirPath, h.Filename)}
		this.Data["json"] = json
	}
	this.ServeJSON()

}

func (this *FileController) UploadPage() {
	this.TplName = "uploadpage.html"
}
