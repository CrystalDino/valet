package models

type Admin struct {
	Id    int64
	Name  string `xorm:"char(8) notnull unique"`
	Cell  string `xorm:"char(16) index notnull unnique"`
	Qq    string `xorm:"char(16)"`
	Email string `xorm:"char(32)"`
	CTime int64  `xorm:"created notnull"`
	MTime int64  `xorm:"updated notnull"`
	DTime int64  `xorm:"deleted"`
}

func (admin *Admin) TableName() string {
	return "admin"
}
