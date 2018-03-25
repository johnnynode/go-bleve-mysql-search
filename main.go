package main

import (
	_ "organ-go-api/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	// "strings"
	// "github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego/plugins/cors"
	// utils "organ-go-api/utils"
)

func init() {
	maxIdle := 50
	maxConn := 50
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDataBase("default", "mysql", "root:gddata#3306@tcp(127.0.0.1:3306)/organ?charset=utf8", maxIdle, maxConn)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 添加过滤器处理
	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods: []string{"PUT", "GET", "DELETE", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	/*
	// 针对所有请求解析token
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		header := ctx.Request.Header // 获取header 类型为map
		auth := header["Authorization"][0] // 解析出 Authorization
		token := strings.Split(auth," ")[1] // 解析出字符串token
		claims, err := utils.JwtDecode(token) // 解析出token中的claims

		if(err == nil) {
			// 创建一个map, 取出所有用户字段 模版： map[sub:zQ6CI4Zq role:16 ip:58.30.138.200 nickname:修改资料检测 oid:gddev exp:1.505444928e+09]
			tokenUser := map[string]interface{}{
				"userId":claims.(jwt.MapClaims)["sub"],
				"role":claims.(jwt.MapClaims)["role"],
				"ip":claims.(jwt.MapClaims)["ip"],
				"nickname":claims.(jwt.MapClaims)["nickname"],
				"oid":claims.(jwt.MapClaims)["oid"],
				"exp":claims.(jwt.MapClaims)["exp"],
			}

			ctx.Input.SetData("tokenUser", tokenUser) // set数据
		}
	})
	*/


	beego.Run()

}
