package controllers

import (
	"encoding/json"
	"organ-go-api/models"
	"strconv"
	// "fmt"
	// "strings"

	"github.com/astaxie/beego"
	// utils "organ-go-api/utils"
)

// 国家相关接口
type CountryController struct {
	beego.Controller
}

// URLMapping ...
func (c *CountryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Country
// @Param	body		body 	models.Country	true		"body for Country content"
// @Success 201 {int} models.Country
// @Failure 403 body is empty
// @router / [post]
func (c *CountryController) Post() {
	var v models.Country
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCountry(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Country by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Country
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CountryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCountryById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description 获取所有国家信息
// @Success 200 {object} models.Country
// @Failure 403
// @router /all [get]
func (c *CountryController) GetAll() {
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

// Put ...
// @Title Put
// @Description update the Country
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Country	true		"body for Country content"
// @Success 200 {object} models.Country
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CountryController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var rb map[string]interface{} // 初始化一个map
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &rb) // 解析出request body
	if err != nil {
		c.Data["json"] = err.Error()
	}else {
		rb["id"] = id // 添加id
		// 更新操作
		if err := models.UpdateCountryById(rb); err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = "ok"
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Country
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CountryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCountry(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
