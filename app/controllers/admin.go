package controllers

import (
	"valet/app/models"

	"github.com/revel/revel"
)

func (c App) Admin() revel.Result {
	c.ViewArgs["bAdmin"] = true
	c.ViewArgs["title"] = "账号管理"
	return c.RenderTemplate("App/admin/admin.html")
}

func (c App) AdminInfo(id int64) revel.Result {
	c.Validation.Required(id)
	if c.Validation.HasErrors() {
		return c.Redirect(App.Admin)
	}
	c.ViewArgs["bSubAdmin"] = true
	c.ViewArgs["bInfoAdmin"] = true
	c.ViewArgs["title"] = "账号详情"
	return c.RenderTemplate("App/admin/info.html")
}

func (c App) AddAdmin() revel.Result {
	c.ViewArgs["bSubAdmin"] = true
	c.ViewArgs["bAddAdmin"] = true
	c.ViewArgs["title"] = "添加账号"
	return c.RenderTemplate("App/admin/add.html")
}

func (c App) UpdateAdmin() revel.Result {
	c.ViewArgs["bSubAdmin"] = true
	c.ViewArgs["bUpdateAdmin"] = true
	c.ViewArgs["title"] = "更新账号"
	return c.RenderTemplate("App/admin/update.html")
}

func (c App) DealAdminAdd(name, cell, password string) revel.Result {
	result := make(map[string]interface{})
	result["error"] = "no"
	c.Validation.Required(name)
	c.Validation.MinSize(name, 4)
	c.Validation.MaxSize(name, 8)

	c.Validation.Required(cell)
	c.Validation.MinSize(cell, 8)
	c.Validation.MaxSize(cell, 12)

	c.Validation.Required(password)
	c.Validation.MinSize(password, 8)
	c.Validation.MaxSize(password, 16)

	if c.Validation.HasErrors() {
		result["error"] = c.Validation.Errors
		return c.RenderJSON(result)
	}

	adm := &models.Admin{
		Name:     name,
		Cell:     cell,
		Password: password,
	}

	id, err := adm.Store()
	if err != nil {
		result["error"] = err.Error()
	}
	result["id"] = id
	return c.RenderJSON(result)
}

func (c App) DealAdminUpdate() revel.Result {
	return c.RenderJSON("")
}
