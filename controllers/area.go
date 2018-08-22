package controllers

import (
	//"encoding/json"
	"loveHome09/models"
	//"time"

	//	"fmt"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	//_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp interface{}) {

	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AreaController) GetAreaInfo() {
	beego.Info("==========/api/v1.0/area get succ!!!=========")

	//把resp返回给前端
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	//函数结束时调用map转json
	defer this.RetData(resp)

	//3.如果redis没有数据，从mysql数据库中获取
	o := orm.NewOrm()
	var areas []models.Area
	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
	if err != nil {
		//返回错误信息
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

		return
	}

	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)

		return
	}
	//最终把resp返回给前端
	resp["data"] = areas

	return
}
