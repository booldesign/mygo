// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"sso/controllers"

	"github.com/astaxie/beego"
)

func init() {
	version := "/v3"
	beego.Router(version+"/token.get", &controllers.TokenController{}, "get,post:GetToken")
	beego.Router(version+"/passport.login", &controllers.PassportController{}, "get,post:Login")
	beego.Router(version+"/passport.users.get", &controllers.UserController{}, "get,post:UserGet")

	beego.Router(version+"/passport.email.send", &controllers.LoginControllor{}, "get,post:EmailSend")
	beego.Router(version+"/passport.sms.send", &controllers.LoginControllor{}, "get,post:SmsSend")
	beego.Router(version+"/passport.verify.code", &controllers.LoginControllor{}, "get,post:VerifyCode")
	beego.Router(version+"/passport.password.reset", &controllers.LoginControllor{}, "get,post:PasswordReset")
	//beego.AddNamespace(ns)

}
