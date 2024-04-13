package config

var Config = config{
	APP: APP{
		PORT:          ":8080",
		SESSIONSECRET: "1234567890",
		JWT:           "1234567890",
		SWAGPATH:      "http://localhost:8080/swagger/docs/index.html#/example",
		BASEURL:       "/api/v1",
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
}
type DB struct {
	NAME     string
	PASSWORD string
	HOST     string
	DB       string
	PORT     string
}
