package config

var Config = config{
	APP: APP{
		PORT:          ":8081",
		SESSIONSECRET: "1234567890",
		JWT:           "1234567890",
		SWAGPATH:      "http://localhost:8081/swagger/docs/index.html#/example",
		BASEURL:       "/cms",
		FILEPATH:      "./uploadFile/",
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
	PORT          string
	SESSIONSECRET string
	JWT           string
	SWAGPATH      string
	BASEURL       string
	FILEPATH      string
}
type DB struct {
	NAME     string
	PASSWORD string
	HOST     string
	DB       string
	PORT     string
}
