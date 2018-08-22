package routers

import (
	"loveHome09/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//请求地域的信息
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreaInfo")

	//请session
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionName;delete:DelSessionName")

	//请index
	beego.Router("/api/v1.0/houses/index", &controllers.HouseeIndexController{}, "get:GetHouseeIndex")

	//请求注册
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")

	//登入请求
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:Login")

	//用户上传头像业务
	//beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UploadAvatar")

}
