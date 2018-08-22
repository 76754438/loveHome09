package controllers

import (
	"loveHome09/models"

	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp interface{}) {

	this.Data["json"] = resp
	this.ServeJSON()
}

//退出
func (this *SessionController) DelSessionName() {

	beego.Info("==========/api/v1.0/session delete  succ!!!=========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	//将session删除
	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")

}

func (this *SessionController) GetSessionName() {
	beego.Info("==========/api/v1.0/session get succ!!!=========")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_SESSIONERR
	resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)

	defer this.RetData(resp)

	name_map := make(map[string]interface{})
	name := this.GetSession("name")

	if name != nil {
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		name_map["name"] = name.(string)
		resp["data"] = name_map
	}

	return

}
