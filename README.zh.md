# 我的应用程序 (my-app)

> [English](./README.md) | [简体中文](./README.zh.md)

“我的应用程序” 是一个不断更新的个人服务集合。

## 初始代码效果预览

> 你可以在[这里](./screenshots/)找到所有的预览截图。

<p align="center">
<img width="49%" src="./screenshots/home.zh.jpg" alt="Home">
<img width="49%" src="./screenshots/settings.zh.jpg" alt="Settings">
</p>

## 技术

| 技术        | 作用                                                  | 来源                                    |
| :---------- | :---------------------------------------------------- | :-------------------------------------- |
| Go          | Backend programming language                          | https://pkg.go.dev/std                  |
| TypeScript  | Frontend programming language                         | https://typescriptlang.org/             |
| Sass/Scss   | Frontend CSS extension language                       | https://sass-lang.com/dart-sass         |
| Vite        | Next Generation Frontend Tooling                      | https://vitejs.dev/                     |
| Vue 3       | Progressive JavaScript Framework                      | https://vuejs.org/                      |
| Wails v2    | Build cross-platform desktop applications using Go    | https://wails.io/                       |
| UPX         | Ultimate packer for executables                       | https://upx.github.io/                  |
| Systray     | A cross platfrom system tray using Go                 | https://github.com/getlantern/systray   |
| Gin         | A HTTP web framework using Go                         | https://gin-gonic.com/                  |
| Gin Swagger | Gin middleware for API documentation with Swagger 2.0 | https://github.com/swaggo/gin-swagger   |
| Swaggo      | Converts Go annotations to Swagger Documentation 2.0  | https://github.com/swaggo/swag          |
| Air         | Live reload and test for API service                  | https://github.com/cosmtrek/air         |
| GORM        | ORM library for Go                                    | https://gorm.io/                        |
| SQLite      | GORM sqlite driver                                    | https://github.com/go-gorm/sqlite       |
| SVG To Font | Generator of fonts from SVG icons                     | https://github.com/jaywcjlove/svgtofont |

## 安装

安装并准备好在 windows 10/11 中的开发环境

- Git v2.37+, https://git-scm.com/
- Go v1.19+, https://go.dev/
- Node.js v16+, https://nodejs.org/
- PNPM v7+, https://pnpm.io/
- VS Code v1.71+ with GCC, https://code.visualstudio.com/docs/cpp/config-mingw
- WebView2 v104+, https://developer.microsoft.com/en-us/microsoft-edge/webview2/

> 安装 VS Code 推荐插件，在插件搜索栏中输入 `@recommended` 查看推荐插件。

> 为 _CGO_ 包做准备，执行命令 `go env -w CGO_ENABLED=1` 。

> 安装项目依赖，在项目根目录中，执行命令 `pnpm install` 。

> 想要更新依赖？在项目根目录中，执行命令 `pnpm update:dependencies` 。

## NPM 脚本

```shell
$ pnpm wails:dev # 开发模式运行（带 wails，tray 和 web）
$ pnpm wails:build # 编译导出 wails 可执行文件
$ pnpm upx:compress # 压缩由 `wails:build` 导出的 wails 可执行文件
$ pnpm air:dev # 单独测试 API 服务
$ pnpm swag:docs # 生成或更新 swagger 文档
$ pnpm docs:dev # 单独测试 vitepress 文档
$ pnpm docs:build # 生成或更新 vitepress 文档
$ pnpm icons:build # 生成 iconfont (frontend/packages/icons)
$ pnpm design:build # 生成组件库 (frontend/packages/design)
$ pnpm <[install|preinstall]:[task]> # 安装项目依赖时会触发这些 install/preinstall 脚本
```

## 目录结构

```yaml
.
├── .tools # 自动生成，开发工具及 CLIs
├── .vscode # VS Code 扩展
├── air # air 热重载工具相关源文件
│   ├── bin # 自动生成，可以尝试 `air:dev` 脚本
│   ├── .air.toml # air 配置
│   └── main.go # 单独运行 Web 服务（不带 wails 和 tray）
├── backend # 后端相关源文件
│   ├── app # app 模块，业务层
│   │   ├── config # 从数据库加载应用设置，全局资源
│   |   │   ├── config.go # 应用设置，与状态和数据操作相关的函数
│   |   │   ├── config.web.go # Web 设置
│   |   │   └── env.go # 加载系统环境变量
│   │   ├── i18n # 后端 i18n 语言源文件，全局资源
│   │   ├── logger # 定义后端 loggers
│   │   ├── app.go # 存放全局状态及资源
│   ├── database # database 模块，持久层
│   │   ├── option # 定义 options 表，用于储存应用设置
│   │   └── database.go # 初始化 database
│   ├── pkg # pkg 模块，横切层
│   │   └── utils # 辅助函数/方法
│   ├── service # service 模块，业务层
│   │   ├── service.go # 初始化提供给服务
│   │   └── settings.go # 应用设置相关函数/方法
│   ├── tray # tray 模块，表示层
│   │   ├── icons # 托盘图标
│   │   ├── menus # 托盘菜单类型
│   │   ├── expose.go # 暴露给 wails 前端的函数
│   │   ├── listen.go # 监听托盘菜单项点击事件
│   │   ├── ready.go # 定义 onReady()，初始化系统托盘菜单
│   │   ├── refresh.go # 定义语言切换和鼠标放置托盘图标时出现提示的刷新函数
│   │   └── tray.go # 系统托盘入口
│   └── web # web 模块，服务层
│       ├── api # API 服务
│       ├── auth # Auth 服务
│       ├── middleware # 自定义的 gin 中间件
│       ├── static # 静态网站及资源
│       │   ├── certs # 自动生成的 localhost TLS 证书, 可以尝试 `install:certs` 脚本
│       │   ├── swagger # 自动生成，可以尝试 `swag:docs` 脚本
│       │   ├── static.go # 管理静态资源
│       │   └── # ...
│       ├── air.go # 包含提供给 air 的特殊方法，用于单独运行 API 服务
│       ├── router.go # swaggo 生成文档入口, 管理 API 路由
│       └── web.go # web 服务入口
├── build # 编译 wails 时用到的资源...
│   ├── bin # 自动生成，可以尝试 `wails:dev` 或 `wails:build` 脚本
│   │   ├── Assets # wails 前端静态资源，自动生成，可以尝试 `wails:build` 脚本
│   │   ├── Docs # vitepress 文档静态资源，自动生成，可以尝试 `docs:build` 脚本
│   │   ├── UserData # webview2 用户数据，应用运行时自动生成
│   │   └── # ...
│   └── # ...
├── diagrams # 4+1 视图模型相关图解
├── docs # vitepress 文档源文件
├── frontend # 前端相关源文件（用 PNPM 管理 workspace）
│   ├── packages # 前端组件、图标字体等等
│   ├── src # wails 前端源文件
│   │   ├── vite-env.d.ts # put go struct associated types into it
│   │   └── # ...
│   └── # ...
├── screenshots # README 应用预览截图
├── main.go # wails main 入口, 表示层
├── wails_life_cycle.go # wails 生命周期
├── wails.json # wails CLI 配置
└── # ...
```

## 开发视图

![Package Diagram](./diagrams/package.png)
