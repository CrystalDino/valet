package controllers

import (
	"valet/app/models"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Manage struct {
	*revel.Controller
}

func (c Manage) Login() revel.Result {
	return c.Render()
}

func (c Manage) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Manage.Login)
}

func (c Manage) CheckLogin(username, password string, remember bool) revel.Result {
	c.Validation.Required(username).Message("必须填写用户名")
	c.Validation.MinSize(username, 5).Message("用户名长度不小于5位字符")
	c.Validation.MaxSize(username, 32).Message("用户名长度不大于30位字符")
	c.Validation.Required(password).Message("必须填写登录密码")
	c.Validation.MinSize(password, 6).Message("密码长度不低于6位字符")
	c.Validation.MaxSize(password, 16).Message("密码长度不大于16位字符")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Manage.Login)
	}
	user := models.GetUserByName(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(App.Index)
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(Manage.Login)
}
