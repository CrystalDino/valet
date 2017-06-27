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
	c.ViewArgs["title"] = "系统信息"
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
		// revel.INFO.Printf("%+v\n", user)
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

func (c App) Admin() revel.Result {
	c.ViewArgs["bAdmin"] = true
	c.ViewArgs["title"] = "账号管理"
	return c.Render()
}

func (c App) Coin() revel.Result {
	c.ViewArgs["bCoin"] = true
	c.ViewArgs["title"] = "币种管理"
	return c.Render()
}

func (c App) User() revel.Result {
	c.ViewArgs["bUser"] = true
	c.ViewArgs["title"] = "用户管理"
	return c.Render()
}

func (c App) Recharge() revel.Result {
	c.ViewArgs["bRecharge"] = true
	c.ViewArgs["title"] = "充值管理"
	return c.Render()
}

func (c App) Withdraw() revel.Result {
	c.ViewArgs["bWithdraw"] = true
	c.ViewArgs["title"] = "提现管理"
	return c.Render()
}

func (c App) CoinIn() revel.Result {
	c.ViewArgs["bCoinIn"] = true
	c.ViewArgs["title"] = "充币管理"
	return c.Render()
}

func (c App) CoinOut() revel.Result {
	c.ViewArgs["bCoinOut"] = true
	c.ViewArgs["title"] = "提币管理"
	return c.Render()
}
