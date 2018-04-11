// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"yoo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/entity_rol",
			beego.NSInclude(
				&controllers.EntityRolController{},
			),
		),

		beego.NSNamespace("/ipc",
			beego.NSInclude(
				&controllers.IpcController{},
			),
		),

		beego.NSNamespace("/entity",
			beego.NSInclude(
				&controllers.EntityController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
