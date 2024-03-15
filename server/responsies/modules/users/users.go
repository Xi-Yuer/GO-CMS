package usersResponsiesModules

import "time"

type SingleUserResponse struct {
	ID         string     `json:"id"`
	Account    string     `json:"account"` // 账号
	Nickname   string     `json:"nickname"`
	Avatar     string     `json:"avatar"`
	Gender     string     `json:"gender"`
	Role       string     `json:"role"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	DeleteTime *time.Time `json:"deleteTime"`
	Status     string     `json:"status"`
}

type QueryUsersParams struct {
	Limit      string `json:"limit" binding:"required"`
	Offset     string `json:"offset" binding:"required"`
	Nickname   string `json:"nickname"`
	Role       string `json:"role"`
	Gender     string `json:"gender"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type CreateUserParams struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserParams struct {
	ID         string `json:"id" binding:"required"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
