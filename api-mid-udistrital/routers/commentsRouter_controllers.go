package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/beego/beego/api-mid-udistrital/controllers:Contactos"] = append(beego.GlobalControllerRouter["github.com/beego/beego/api-mid-udistrital/controllers:Contactos"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
