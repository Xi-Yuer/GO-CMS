// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/bindRoles": {
            "post": {
                "description": "给用户分配角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "给用户分配角色",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/captcha": {
            "get": {
                "description": "获取验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "获取验证码",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "登录",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pages": {
            "post": {
                "description": "创建页面",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "页面管理"
                ],
                "summary": "创建页面",
                "responses": {}
            }
        },
        "/pages/{id}": {
            "delete": {
                "description": "删除页面",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "页面管理"
                ],
                "summary": "删除页面",
                "responses": {}
            }
        },
        "/roles": {
            "get": {
                "description": "获取角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "获取角色",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "创建角色",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/roles/{id}": {
            "put": {
                "description": "更新角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "更新角色",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除角色",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "删除角色",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "获取系统中所有用户的列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码，默认为1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示的用户数量，默认为10",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "创建新用户并将其添加到系统中",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建新用户",
                "responses": {}
            }
        },
        "/users/search": {
            "get": {
                "description": "根据查询参数查询用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "查询用户",
                "responses": {}
            }
        },
        "/users/{id}": {
            "get": {
                "description": "根据ID获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "根据ID获取用户信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "更新现有用户的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "根据用户ID删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
