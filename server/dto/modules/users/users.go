package usersResponsiesModules

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type QueryUsersParams struct {
	Limit        int    `form:"limit"`
	Offset       int    `form:"offset"`
	ID           string `form:"id"`
	Nickname     string `form:"nickname"`
	Account      string `form:"account"`
	Gender       string `form:"gender"`
	DepartmentID string `form:"departmentID"`
	Status       string `form:"status"`
	CreateTime   string `form:"createTime"`
	UpdateTime   string `form:"updateTime"`
	StartTime    string `form:"startTime"`
	EndTime      string `form:"endTime"`
}

type CreateUserParams struct {
	Account      string   `form:"account" json:"account" binding:"required"`
	Nickname     string   `form:"nickname" json:"nickname" binding:"required"`
	Password     string   `form:"password" json:"password" binding:"required"`
	RoleID       []string `form:"roleID" json:"rolesID" binding:"required"`
	DepartmentID string   `form:"departmentID" json:"departmentID" binding:"required"`
}

type UpdateUserParams struct {
	Nickname     string   `form:"nickname"`
	Password     string   `form:"password"`
	RolesID      []string `form:"rolesID"`
	Status       string   `form:"status"`
	DepartmentID string   `form:"departmentID"`
}

type Page struct {
	Limit  *int `form:"limit" json:"limit" binding:"required"`
	Offset *int `form:"offset" json:"offset" binding:"required"`
}

type SingleUserResponse struct {
	ID           string     `json:"id"`
	Account      string     `json:"account"` // 账号
	Nickname     string     `json:"nickname"`
	IsAdmin      int        `json:"isAdmin"`
	Avatar       *string    `json:"avatar"`
	RolesID      []string   `json:"rolesID"`
	InterfaceDic []string   `json:"interfaceDic"`
	DepartmentID *string    `json:"departmentID"`
	CreateTime   *time.Time `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
	Status       string     `json:"status"`
}
type SingleUserByRoleIDResponse struct {
	ID           string     `json:"id"`
	Account      string     `json:"account"` // 账号
	Nickname     string     `json:"nickname"`
	Avatar       *string    `json:"avatar"`
	IsAdmin      int        `json:"isAdmin"`
	RolesID      []string   `json:"rolesID"`
	DepartmentID *string    `json:"departmentID"`
	CreateTime   *time.Time `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime"`
	Status       string     `json:"status"`
}

type SingleUserResponseHasPassword struct {
	SingleUserResponse
	Password string `json:"password"`
}

type JWTPayload struct {
	ID           string   `json:"id"`
	Account      string   `json:"account"`
	IsAdmin      int      `json:"isAdmin"`
	RoleID       []string `json:"roleId"`
	InterfaceDic []string `json:"interfaceDic"`
	DepartmentID *string  `json:"departmentId"`
	jwt.RegisteredClaims
}

type QueryUsersByRoleIDs struct {
	RoleIDs []string `form:"roleIDs"`
}
