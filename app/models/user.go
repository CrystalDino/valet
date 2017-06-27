package models

type User struct {
	Id        int64
	Name      string
	Cell      string
	Password  string
	Transcode string
	Stat      int8
	CTime     int64
	MTime     int64
	DTime     int64
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
