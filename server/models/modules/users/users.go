package userModelsModule

import "time"

var UserModel = &User{}

type User struct {
	ID         string     `json:"id"`
	Account    string     `json:"account"`  // 账号
	Password   string     `json:"password"` // 密码
	Nickname   string     `json:"nickname"`
	Avatar     string     `json:"avatar"`
	Gender     string     `json:"gender"`
	Role       string     `json:"role"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	DeleteTime *time.Time `json:"deleteTime"`
	Status     string     `json:"status"`
}
