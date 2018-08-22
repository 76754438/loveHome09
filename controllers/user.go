package controllers

import (
	"loveHome09/models"

	"encoding/json"

	"github.com/astaxie/beego"

	//"path"

	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp interface{}) {

	this.Data["json"] = resp
	this.ServeJSON()
}

//1.在conf中添加copyrequestbody = true
//2.写获取用户输入的用户名和密码
func (this *UserController) Reg() {
	beego.Info("==========/api/v1.0/user get succ!!!=========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
	//存储前端的信息
	var regRequestMap = make(map[string]interface{})
	//1.得到前端请求的数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	beego.Info("mobile=", regRequestMap["mobile"])     //手机号
	beego.Info("password=", regRequestMap["password"]) //密码
	beego.Info("sms_code=", regRequestMap["sms_code"]) //短信验证码

	//2.判断数据的合理性
	if regRequestMap["mobile"] == "" || regRequestMap["password"] == "" || regRequestMap["sms_code"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3.将数据存储到mysql的user表
	user := models.User{}
	user.Mobile = regRequestMap["mobile"].(string)

	user.Password_hash = regRequestMap["password"].(string)

	user.Name = regRequestMap["mobile"].(string)

	//操作把数据入数据库
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert err=", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	}

	//显示注册成功
	beego.Info("reg succ ! id=", id) //如果id=1，表示成功把数据存储到数据库中

	//4.将当前用户的信息存储到session中
	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", id)
	this.SetSession("mobile", user.Mobile)

	return

}

//登入业务
func (this *UserController) Login() {
	beego.Info("==========/api/v1.0/user Login succ!!!=========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	var loginRequestMap = make(map[string]interface{})

	//1.得到客户端请求的json数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginRequestMap)
	beego.Info("mobile=", loginRequestMap["mobile"])     //手机号
	beego.Info("password=", loginRequestMap["password"]) //密码

	//2.判断数据的合理性
	if loginRequestMap["mobile"] == "" || loginRequestMap["password"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3.查询数据库得到user
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if err := qs.Filter("mobile", loginRequestMap["mobile"]).One(&user); err != nil {
		//查询失败,无数据
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	//4.对比密码
	if user.Password_hash != loginRequestMap["password"].(string) {
		resp["errno"] = models.RECODE_PWDERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		return
	}

	beego.Info("==== Login succ =======user.Name:", user.Name)

	//5.将当前用户的信息存储到session中
	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)

}
