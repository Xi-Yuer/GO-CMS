package constant

import "github.com/Xi-Yuer/cms/config"

var AUTHORIZATION = `Authorization`

var JWTPAYLOAD = `JwtPayload`

var PermissionMap = map[string]string{
	// 用户管理
	"^POST:" + config.Config.APP.BASEURL + "/users/query$":           `POST:/users/query`,
	"^POST:" + config.Config.APP.BASEURL + "/users$":                 `POST:/users`,
	"^GET:" + config.Config.APP.BASEURL + "/users/role/\\d+$":        `GET:/users/role/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/users/\\d+$":             `GET:/users/:id`,
	"^PATCH:" + config.Config.APP.BASEURL + "/users/\\d+$":           `PATCH:/users/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/users/\\d+$":          `DELETE:/users/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/users/export$":          `POST:/users/export`,
	"^POST:" + config.Config.APP.BASEURL + "/users/query/role/\\d+$": `POST:/users/query/role/:id`,

	// 角色管理
	"^POST:" + config.Config.APP.BASEURL + "/roles/query$":     `POST:/roles/query`,
	"^GET:" + config.Config.APP.BASEURL + "/users/role/\\d+$":  `GET:/users/role/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/roles$":           `POST:/roles`,
	"^PATCH:" + config.Config.APP.BASEURL + "/roles/\\d+$":     `PATCH:/roles/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/roles/\\d+$":    `DELETE:/roles/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/roles/export$":    `POST:/roles/export`,
	"^POST:" + config.Config.APP.BASEURL + "/roles/bindUser$":  `POST:/roles/bindUser`,
	"^POST:" + config.Config.APP.BASEURL + "/roles/deBindUser": `POST:/roles/deBindUser`,

	// 部门管理
	"^GET:" + config.Config.APP.BASEURL + "/department$":         `GET:/department`,
	"^POST:" + config.Config.APP.BASEURL + "/department$":        `POST:/department`,
	"^PATCH:" + config.Config.APP.BASEURL + "/department/\\d+$":  `PATCH:/department/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/department/\\d+$": `DELETE:/department/:id`,

	// 菜单管理
	"^GET:" + config.Config.APP.BASEURL + "/pages$":           `GET:/pages`,
	"^POST:" + config.Config.APP.BASEURL + "/pages$":          `POST:/pages`,
	"^PATCH:" + config.Config.APP.BASEURL + "/pages/\\d+$":    `PATCH:/pages/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/pages/\\d+$":   `DELETE:/pages/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/pages$":           `GET:/pages`,
	"^GET:" + config.Config.APP.BASEURL + "/pages/user":       `GET:/pages/user`,
	"^GET:" + config.Config.APP.BASEURL + "/pages/role:\\d+$": `GET:/pages/role/id`,

	// 接口管理
	"^POST:" + config.Config.APP.BASEURL + "/interface$":          `POST:/interface`,
	"^GET:" + config.Config.APP.BASEURL + "/interface/page/\\d+$": `GET:/interface/page/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/interface/role/\\d+$": `GET:/interface/role/:id`,
	"^DELETE:" + config.Config.APP.BASEURL + "/interface/\\d+$":   `DELETE:/interface/:id`,
	"^PATCH:" + config.Config.APP.BASEURL + "/interface/\\d+$":    `PATCH:/interface/:id`,

	// 日志管理
	"^GET:" + config.Config.APP.BASEURL + "/log/commits$": `GET:/log/commits`,
	"^GET:" + config.Config.APP.BASEURL + "/log/system$":  `GET:/log/system`,

	// 系统管理
	"^GET:" + config.Config.APP.BASEURL + "/system$": `GET:/system`,

	// 定时任务
	"^POST:" + config.Config.APP.BASEURL + "/timeTask/start/\\d+$":   `POST:/timeTask/start/:id`,
	"^POST:" + config.Config.APP.BASEURL + "/timeTask/stop/\\d+$":    `POST:/timeTask/stop/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/timeTask$":               `GET:/timeTask`,
	"^PATCH:" + config.Config.APP.BASEURL + "/timeTask/update/\\d+$": `PATCH:/timeTask/update/:id`,

	// 文件上传下载
	"^DELETE:" + config.Config.APP.BASEURL + "/upload/del/.+$": `DELETE:/upload/del/:id`,
	"^GET:" + config.Config.APP.BASEURL + "/upload$":           `GET:/upload`,
	"^POST:" + config.Config.APP.BASEURL + "/upload/check$":    `POST:/upload/check`,
	"^POST:" + config.Config.APP.BASEURL + "/upload$":          `POST:/upload`,
	"^POST:" + config.Config.APP.BASEURL + "/upload/finish$":   `POST:/upload/finish`,
	// 是否可以通过 Ajax 下载文件
	"^POST:" + config.Config.APP.BASEURL + "/upload/download/.+$": `POST:/upload/download/:id`,
	// 是否可以通过 a 连接下载文件
	"^GET:" + config.Config.APP.BASEURL + "/upload/download/aHref/.+$": `GET:/upload/download/aHref/:id`,

	// 代码生成器
	"^POST:" + config.Config.APP.BASEURL + "/template":           `POST:/template`,
	"^POST:" + config.Config.APP.BASEURL + "/template/download$": `POST:/template/download`,
}
