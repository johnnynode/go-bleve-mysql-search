package controllers

import (
	"encoding/json"
	// "errors"
	"organ-go-api/models"
	"strconv"
	// "strings"

	"github.com/astaxie/beego"
	"fmt"
	"organ-go-api/utils"
)

// 机构增删改查
type OrganController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrganController) URLMapping() {
	c.Mapping("PostNewOrgan", c.PostNewOrgan)
	c.Mapping("GetOrganByUuid", c.GetOrganByUuid)
	c.Mapping("GetOrganList", c.GetOrganList)
	c.Mapping("UpdateOrgan", c.UpdateOrgan)
	c.Mapping("DeleteOrganByUuid", c.DeleteOrganByUuid)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Post Organ Add
// @Description 添加新机构
// @Param	body		body 	models.Organ	true		"body for Organ content"
// @Success 201 {object} models.Organ
// @Failure 403 body is empty
// @router /add [post]
func (c *OrganController) PostNewOrgan() {
	var o models.Organ
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &o); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if _, err := models.AddOrgan(&o); err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = o
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description 通过uuid查询机构详情
// @Param	uuid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Organ
// @Failure 403 :id is empty
// @router /detail/:uuid [get]
func (c *OrganController) GetOrganByUuid() {
	uuid := c.Ctx.Input.Param(":uuid")
	OrganFull, err := models.GetOrganByUuid(uuid) // 通过uuid查出机构表信息
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		OrganAliasList, err := models.GetOrganAliasByUuid(uuid) // 通过uuid查处机构别名信息
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			OrganFull.OrganAlias = OrganAliasList
			c.Data["json"] = OrganFull
		}
	}
	c.ServeJSON()
}

// GetOrganList ...
// @Title Get Organ List
// @Description 通过机构 oid, country_code, uuid 获取相关的机构列表
// @Param	keyword	query	string	true	"oid, country_code, uuid"
// @Param	num	  path	  integer	true	"number"
// @Success 200 {object} models.Organ
// @Failure 403
// @router /getorganlist/:num [get]
func (c *OrganController) GetOrganList() {
	keyword := c.GetString(":keyword") // 获取关键字
	numStr := c.Ctx.Input.Param(":num") // 获取num
	num, _ := strconv.Atoi(numStr)      // 将num转换为数字
	pageNumber := int64(num) // 当前页数
	total := models.GetTotalByKeyword(keyword) // 获取通过关键字查询获取的总数
	list, err := models.GetOrganList(keyword, int64(num)) // 获取查询结果的集合

	var pages int64 = utils.GetPages(total, 20) // 计算总页数
	// 定义一个Page对象
	p := utils.Page{
		List:                list,
		Total:               total,
		Limit:               20, // 初始化Limit
		Pages:               pages,
		PageNumber:          pageNumber,
		NavigatePageNumbers: utils.CalcNavigatePageNumbers(pageNumber, pages),
		FirstPage:           utils.IsFirstPage(pageNumber),
		LastPage:            utils.IsLastPage(pageNumber, pages),
	}

	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = p
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description 更新机构字段，id一定要传
// @Param	body		body 	models.Organ	true		"body for Organ content"
// @Success 200 {object} models.Organ
// @Failure 403 :id is not int
// @router /update [put]
func (c *OrganController) UpdateOrgan() {
	var v models.Organ
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if err := models.UpdateOrgan(&v); err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = "ok"
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description 通过uuid删除机构
// @Param	uuid		path 	string	true		"The uuid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uuid is empty
// @router /delete/:uuid [delete]
func (c *OrganController) DeleteOrganByUuid() {
	uuid := c.Ctx.Input.Param(":uuid")
	fmt.Println("uuid: ",uuid)
	if err := models.DeleteOrgan(uuid); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description 获取所有国家信息
// @Success 200 {object} models.Country
// @Failure 403
// @router /getcountry [get]
func (c *OrganController) GetAll() {
	// tokenUser := c.Ctx.Input.GetData("tokenUser").(map[string]interface{}) // 获取token 用户
	// fmt.Println(tokenUser["userId"]) // 成功获取用户 userId
	l, err := models.GetAllCountry()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
