package controllers

import (
	"encoding/json"
	"organ-go-api/models"
	"strconv"
	"github.com/astaxie/beego"
)

// OrganAliasController operations for OrganAlias
type OrganAliasController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrganAliasController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create OrganAlias
// @Param	body		body 	models.OrganAlias	true		"body for OrganAlias content"
// @Success 201 {int} models.OrganAlias
// @Failure 403 body is empty
// @router /add [post]
func (c *OrganAliasController) Post() {
	var v models.OrganAlias
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if _, err := models.AddOrganAlias(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get OrganAlias by id
// @Param	uuid		path 	string	true		"The key to find"
// @Success 200 {object} models.OrganAlias
// @Failure 403 :uuid is empty
// @router /:uuid [get]
func (c *OrganAliasController) GetOne() {
	uuid := c.Ctx.Input.Param(":uuid")
	v, err := models.GetOrganAliasByUuid(uuid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrganAlias
// @Param	body		body 	models.OrganAlias	true		"body for OrganAlias content"
// @Success 200 {object} models.OrganAlias
// @Failure 403
// @router /update [put]
func (c *OrganAliasController) Put() {
	var v models.OrganAlias
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
	} else {
		if err := models.UpdateOrganAliasById(&v); err == nil {
			c.Data["json"] = "ok"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the OrganAlias
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrganAliasController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrganAlias(id); err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "ok"
	}
	c.ServeJSON()
}
