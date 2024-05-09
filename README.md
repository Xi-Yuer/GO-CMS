# 基于RBAC的权限管理系统

## 项目介绍
基于RBAC的权限管理系统，使用Go + Gin框架，Mysql作为数据库，前端采用 React + Antd 实现。支持国际化，主题色切换，权限控制，菜单管理，用户管理，角色管理，数据字典管理，日志管理，系统监控，代码生成器，接口文档等功能。

## 项目运行

1. 克隆项目到本地
2. 运行项目设置环境变量
    ```text
      {
        PORT:          ":8080",
        SESSIONSECRET: "sessionSecrete",
        JWT:           "jwtSecrete",
        SWAGPATH:      "http://localhost:8080/swagger/docs/index.html#/example",
        BASEURL:       "/api/v1",
        FILEPATH:      "./uploadFile/",
      }
    ```
3. 访问 http://localhost:8080/api/v1
4. 用户名：admin 密码：123

## 项目部分截图
![登录](./static/login.png)

![首页](./static/dashboard.png)

![系统管理](./static/system.png)