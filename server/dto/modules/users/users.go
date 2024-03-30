package usersResponsiesModules

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type QueryUsersParams struct {
	Limit      *string `form:"limit" binding:"required"`
	Offset     *string `form:"offset" binding:"required"`
	ID         string  `form:"id"`
	Nickname   string  `form:"nickname"`
	Account    string  `form:"account"`
	Role       string  `form:"role"`
	Gender     string  `form:"gender"`
	Status     string  `form:"status"`
	CreateTime string  `form:"createTime"`
	UpdateTime string  `form:"updateTime"`
	StartTime  string  `form:"startTime"`
	EndTime    string  `form:"endTime"`
}

type CreateUserParams struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserParams struct {
	Nickname string `form:"nickname"`
	Password string `form:"password"`
}

type Page struct {
	Limit  *int `form:"limit" binding:"required"`
	Offset *int `form:"offset" binding:"required"`
}

type SingleUserResponse struct {
	ID         string     `json:"id"`
	Account    string     `json:"account"` // 账号
	Nickname   string     `json:"nickname"`
	Avatar     *string    `json:"avatar"`
	RolesID    []string   `json:"rolesID"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	Status     *string    `json:"status"`
}

type SingleUserResponseHasPassword struct {
	SingleUserResponse
	Password string `json:"password"`
}

type JWTPayload struct {
	ID               string   `json:"id"`
	NickName         string   `json:"nickName"`
	Role             string   `json:"role"`
	PagePermission   []string `json:"pagePermission"`
	ButtonPermission []string `json:"buttonPermission"`
	jwt.RegisteredClaims
}
