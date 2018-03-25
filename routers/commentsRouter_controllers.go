package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:CountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "ActiveIndex",
			Router: `/active/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "CreateIndex",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "DeleteIndex",
			Router: `/delete/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "ShowOneIndex",
			Router: `/detail/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "ListIndex",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "StopIndex",
			Router: `/stop`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganAliasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/update`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "PostNewOrgan",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "DeleteOrganByUuid",
			Router: `/delete/:uuid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "GetOrganByUuid",
			Router: `/detail/:uuid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getcountry`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "GetOrganList",
			Router: `/getorganlist/:num`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:OrganController"],
		beego.ControllerComments{
			Method: "UpdateOrgan",
			Router: `/update`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"],
		beego.ControllerComments{
			Method: "ClientSearchOrgan",
			Router: `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"],
		beego.ControllerComments{
			Method: "CommonSearch",
			Router: `/search`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"] = append(beego.GlobalControllerRouter["organ-go-api/controllers:SearchController"],
		beego.ControllerComments{
			Method: "ClientSearchSchool",
			Router: `/searchschool`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
