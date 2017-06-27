package models

type User struct {
	Id        int64
	Name      string `xorm:"char(32) notnull"`
	Cell      string `xorm:"char(16) notnull index unique"`
	Password  string `xorm:"char(128) notnull"`
	Transcode string `xorm:"char(128)"`
	Email     string `xorm:"char(32)"`
	Stat      int8
	CTime     int64 `xorm:"created notnull"`
	MTime     int64 `xorm:"updated notnull"`
	DTime     int64 `xorm:"deleted"`
}

func (user *User) TableName() string {
	return "user"
}

func GetUserByName(username string) (user *User) {
	//test data
	user = &User{
		Id:        10234,
		Name:      "dino",
		Cell:      "123123123",
		Password:  "$2a$10$UM9gvkZU5Nf5qllQsGiIAeixDf6mH95/mbML1qWwJmvprU5sHabD2",
		Transcode: "654321",
		Stat:      1,
		CTime:     111111,
		MTime:     222222,
		DTime:     333333,
	}
	return
}
