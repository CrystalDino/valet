package controllers

import (
	"valet/app/models"

	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(models.InitDB)
	revel.InterceptMethod(App.SetUser, revel.BEFORE)
	revel.InterceptMethod(App.CheckUser, revel.BEFORE)
}
