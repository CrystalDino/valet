package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id       int64
	Name     string `xorm:"char(8) notnull unique"`
	Cell     string `xorm:"char(16) index notnull unnique"`
	Password string `xorm:"char(128) notnull"`
	Qq       string `xorm:"char(16)"`
	Email    string `xorm:"char(32)"`
	Icon     string `xorm:"char(128)"`
	Info     string `xorm:"varchar(256)"`
	Stat     int8   //1：新注册，未修改密码 2：正常用户，已修改密码 3：禁用
	CTime    int64  `xorm:"created notnull"`
	MTime    int64  `xorm:"updated notnull"`
	DTime    int64  `xorm:"deleted"`
}

func (admin *Admin) TableName() string {
	return "admin"
}

func (admin *Admin) Store() (id int64, err error) {
	pas, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	admin.Stat = 1 //new admin
	admin.Password = string(pas)
	if id, err = engine.InsertOne(admin); err != nil {
		return -1, err
	}
	if id != 1 {
		return -1, errors.New("db error: insert count not one")
	}
	var adm = &Admin{Name: admin.Name}
	has, err := engine.Select("id").Get(adm)
	if err != nil {
		return -1, err
	}
	if !has {
		return -1, errors.New("create admin account " + admin.Name + " failed")
	}
	println("password:", adm.Password)
	id = adm.Id
	return
}
