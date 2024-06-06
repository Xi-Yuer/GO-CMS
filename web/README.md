# React + TypeScript + Vite

```bash
#安装依赖
pnpm install
#启动开发服务器
pnpm run dev
#构建生产版本
pnpm run build
#启动生产服务器
pnpm run preview
```

### 技术栈

- React
  - https://react.docschina.org/
- TypeScript
  - https://www.typescriptlang.org/
- Vite
  - https://cn.vitejs.dev/
- React-Router
  - https://reactrouter.com/en/main 
- Ant Design
  - https://ant.design/index-cn
- TailwindCss
  - https://tailwindcss.com/
- ReactRTK
  - https://redux.js.org/rtk/overview
- ReactI18n
  - https://react.i18next.com/
- Echarts
  - https://echarts.apache.org/zh/index.html

### 项目结构

```text
src
    ├── LayOut                                      # 布局
    ├── assets                                      # 静态资源
    │   ├── imag
    │   └── svg
    ├── components                                  # 全局组件
    │   ├── AppBreadcrumb
    │   ├── AppHeaderTab
    │   ├── AppUploads
    │   ├── Auth
    │   ├── Card
    │   ├── CodeEdit
    │   ├── Footer
    │   ├── Icon
    │   ├── IconSelect
    │   ├── Iframe
    │   ├── Theme
    │   └── Translate
    ├── constant                                    # 常量
    ├── hooks                                       # 全局hooks
    ├── local                                       # 国际化
    │   ├── en
    │   └── zh
    ├── pages                                       # 页面
    │   ├── About
    │   ├── Dashboard
    │   ├── File
    │   │   └── Upload
    │   │       └── hooks
    │   ├── Login
    │   ├── Monitor
    │   │   ├── Logs
    │   │   └── TimeTask
    │   ├── NotFond
    │   ├── System
    │   │   ├── Department
    │   │   ├── Menu
    │   │   ├── Role
    │   │   └── User
    │   └── SystemUtils
    │       └── CodeGenerate
    ├── router                                      # 路由
    ├── service                                     # 接口
    │   ├── api
    │   │   ├── department
    │   │   ├── file
    │   │   ├── interface
    │   │   ├── login
    │   │   ├── logs
    │   │   ├── pages
    │   │   ├── roles
    │   │   ├── system
    │   │   ├── template
    │   │   ├── timeTask
    │   │   └── users
    │   └── request
    │       └── lib
    ├── store                                        # 全局store
    │   ├── UIStore
    │   ├── UploadStore
    │   └── UserStore
    ├── styles                                       # 全局样式
    │   ├── common
    │   └── global
    ├── theme                                        # 主题
    ├── types                                        # 全局类型
    │   └── menus
    └── utils                                        # 全局工具类
        ├── builder
        ├── cache
        ├── echarts
        ├── event
        ├── format
        ├── message
        ├── monaco
        └── sleep
```