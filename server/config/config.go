package config

var Config = config{
	APP: APP{
		PORT: ":8080",
	},
	DB: DB{
		NAME:     "root",
		PASSWORD: "2214380963Wx!!",
		HOST:     "localhost",
		DB:       "cms",
		PORT:     "3306",
	},
}

type config struct {
	APP
	DB
}

type APP struct {
	PORT string
}
type DB struct {
	NAME     string
	PASSWORD string
	HOST     string
	DB       string
	PORT     string
}
