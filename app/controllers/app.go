package controllers

import (
	"valet/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.ViewArgs["bHome"] = true
	c.ViewArgs["title"] = "Dashboard"
	return c.Render()
}

func (c App) getUserFromView() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return models.GetUserByName(username)
	}
	return nil
}

func (c App) SetUser() revel.Result {
	if user := c.getUserFromView(); user != nil {
		c.ViewArgs["user"] = user
		revel.INFO.Printf("%+v\n", user)
	}
	return nil
}

func (c App) CheckUser() revel.Result {
	if user := c.getUserFromView(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(Manage.Login)
	}
	return nil
}
