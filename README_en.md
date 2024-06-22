# RBAC-Based Permission Management System

[简体中文](README.md) | English

[![GoDoc](https://godoc.org/github.com/Xi-Yuer/GO-CMS?status.svg)](https://godoc.org/github.com/Xi-Yuer/GO-CMS)
![License](https://img.shields.io/github/license/Xi-Yuer/GO-CMS)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/Xi-Yuer/GO-CMS)
![GitHub last commit](https://img.shields.io/github/last-commit/Xi-Yuer/GO-CMS)
![GitHub issues](https://img.shields.io/github/issues/Xi-Yuer/GO-CMS)
![GitHub forks](https://img.shields.io/github/forks/Xi-Yuer/GO-CMS)
![GitHub stars](https://img.shields.io/github/stars/Xi-Yuer/GO-CMS)

## Project Introduction
The RBAC-based permission management system is developed using the Go + Gin framework, with MySQL as the database and React + Antd for the frontend. It supports features such as internationalization, theme color switching, permission control, menu management, user management, role management, data dictionary management, log management, system monitoring, code generator, and API documentation.

## Backend Setup

1. Clone the project to your local machine
   ```text
   https://github.com/Xi-Yuer/GO-CMS.git
   ```
   
2. Install the dependencies
    ```text
    cd server
    go mod tidy
    ```
   
3. Modify the project settings to set environment variables
   server/config/config.go
    ```text
     {
        NAME:     "root",  // Database username
        PASSWORD: "xxxxxx", // Password
        HOST:     "localhost", // Host address
        DB:       "cms", // Database name
        PORT:     "3306", // Port
     }
   ```

4. Create MySQL runtime environment

    1. Create a database named cms, execute the SQL file (location: server/sql/cms_widthData.sql)

5. Run the project
    ```bash
      go run server/main.go
    ```
   


## Frontend Setup

1. Clone the project to your local machine
   ```text
   https://github.com/Xi-Yuer/GO-CMS.git
   ```
   
2. Install the dependencies
    ```text
    cd client
    pnpm install
    ```
   
3. Run the project
    ```bash
      pnpm run dev
    ```
   

## Project Screenshots

![登录](./static/login.png)

![首页](./static/dashboard.png)

![系统管理](./static/system.png)