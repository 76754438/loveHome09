package controllers

import (
	"loveHome09/models"

	"github.com/astaxie/beego"
)

type HouseeIndexController struct {
	beego.Controller
}

func (this *HouseeIndexController) RetData(resp interface{}) {

	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseeIndexController) GetHouseeIndex() {
	beego.Info("==========/api/v1.0/houseeindex get succ!!!=========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

}
