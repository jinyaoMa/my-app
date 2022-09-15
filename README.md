# My App (my-app)

My App is a continuously updated personal service collection.

## Setup

Prepare and install environment for development in Window 10/11?

- Git v2.37+, https://git-scm.com/
- Go v1.19+, https://go.dev/
- Node.js v16+, https://nodejs.org/
- PNPM v7+, https://pnpm.io/
- VS Code v1.71+ with GCC, https://code.visualstudio.com/docs/cpp/config-mingw
- WebView2 v104+, https://developer.microsoft.com/en-us/microsoft-edge/webview2/

> Setup VS Code by installing recommended extensions. To do this, enter in `@recommended` while searching for extensions.

> Run command `go env -w CGO_ENABLED=1` to prepare for _CGO_ enabled packages

> Run command `pnpm install` at project root directory to setup.

## Technologies

| Technology  | Role                                                  | Sources                               |
| :---------- | :---------------------------------------------------- | :------------------------------------ |
| Go          | Backend programming language                          | https://pkg.go.dev/std                |
| TypeScript  | Frontend programming language                         | https://typescriptlang.org/           |
| Vite        | Next Generation Frontend Tooling                      | https://vitejs.dev/                   |
| Vue 3       | Progressive JavaScript Framework                      | https://vuejs.org/                    |
| Wails       | Build cross-platform desktop applications using Go    | https://wails.io/                     |
| UPX         | Ultimate packer for executables                       | https://upx.github.io/                |
| Systray     | A cross platfrom system tray using Go                 | https://github.com/getlantern/systray |
| Gin         | A HTTP web framework using Go                         | https://gin-gonic.com/                |
| Gin Swagger | Gin middleware for API documentation with Swagger 2.0 | https://github.com/swaggo/gin-swagger |
| Swaggo      | Converts Go annotations to Swagger Documentation 2.0  | https://github.com/swaggo/swag        |
| Air         | Live reload and test for API service                  | https://github.com/cosmtrek/air       |
| GORM        | ORM library for Go                                    | https://gorm.io/                      |
| SQLite      | GORM sqlite driver                                    | https://github.com/go-gorm/sqlite     |
