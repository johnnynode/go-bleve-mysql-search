// @APIVersion 1.0.0
// @Title Organ 机构管理
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"organ-go-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/organalias",
			beego.NSInclude(
				&controllers.OrganAliasController{},
			),
		),

		beego.NSNamespace("/country",
			beego.NSInclude(
				&controllers.CountryController{},
			),
		),

		beego.NSNamespace("/organ",
			beego.NSInclude(
				&controllers.OrganController{},
			),
		),

		beego.NSNamespace("/index",
			beego.NSInclude(
				&controllers.IndexController{},
			),
		),

		beego.NSNamespace("/organindexsearch",
			beego.NSInclude(
				&controllers.SearchController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
