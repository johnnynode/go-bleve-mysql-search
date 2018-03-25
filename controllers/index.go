package controllers

import (
	"github.com/astaxie/beego"
	utils "organ-go-api/utils"
	"fmt"
)

// 索引管理相关接口
type IndexController struct {
	beego.Controller
}

// URLMapping ...
func (i *IndexController) URLMapping() {
	i.Mapping("CreateIndex", i.CreateIndex) // 创建索引 post
	i.Mapping("ActiveIndex", i.ActiveIndex) // 激活索引 get
	i.Mapping("ListIndex", i.ListIndex) // 列出所有索引 get
	i.Mapping("ShowOneIndex", i.ShowOneIndex) // 列出单条索引详情 get
	i.Mapping("StopIndex", i.StopIndex) // 停止创建索引 get
	i.Mapping("DeleteIndex", i.DeleteIndex) // 删除索引 delete
}

// 创建索引 post
// @Title CreateIndex 备注这个`@`貌似没用
// @Description 创建索引
// @Success 201
// @Failure 403
// @router /create [post]
func (i *IndexController) CreateIndex() {
	fmt.Println("收到新建索引请求")
	i.Data["json"] = "构建中,请去控制台查看..."
	i.ServeJSON()
	go utils.CreateIndex() // 开始构建
}

// 激活索引 Get
// @Title ActiveIndex
// @Description 激活索引
// @Param  id	 path 	string	 true  	"The key to active"
// @Success 200
// @Failure 403
// @router /active/:id [get]
func (i *IndexController) ActiveIndex() {
	id := i.Ctx.Input.Param(":id")
	fmt.Println("收到激活索引请求", id)
	go utils.ActiveIndex(id)
}

// 列出所有索引 Get
// @Title ListIndex
// @Description 列出所有索引
// @Success 200
// @Failure 403
// @router /list [get]
func (i *IndexController) ListIndex() {
	fmt.Println("收到列出所有索引请求")
	list, err := utils.ListIndex()
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = list
	}
	i.ServeJSON()
}

// 列出单条索引详情 Get
// @Title ShowOneIndex
// @Description 列出单条索引详情
// @Param  id	 path 	string	 true  	"The key to find specific Index"
// @Success 200
// @Failure 403
// @router /detail/:id [get]
func (i *IndexController) ShowOneIndex() {
	id := i.Ctx.Input.Param(":id") // 获取参数
	fmt.Println("收到列出单条索引的详情请求")
	ii, err := utils.ListOneIndex(id)
	if err != nil {
		i.Data["json"] = err.Error()
	} else {
		i.Data["json"] = ii
	}
	i.ServeJSON()
}

// 停止创建索引 Get
// @Title StopIndex
// @Description 停止创建索引
// @Success 200
// @Failure 403
// @router /stop [get]
func (i *IndexController) StopIndex() {
	go utils.StopIndex() // 停止
}

// 删除索引
// @Title DeleteIndex
// @Description 删除索引
// @Param  id	 path 	string	 true  	"The key to delete"
// @Success 200
// @Failure 403
// @router /delete/:id [delete]
func (i *IndexController) DeleteIndex() {
	fmt.Println("收到删除索引请求：")
	id := i.Ctx.Input.Param(":id")
	fmt.Println("i: ",id)
	// 定义info规范："0" => 没有一条索引， "1" => 删除成功 , "2" => 索引不存在
	info := utils.DeleteIndex(id)
	i.Data["json"] = info
	i.ServeJSON()
}