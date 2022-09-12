# My App (my-app)

Prepare environment for development?

- Go v1.19+, https://go.dev/
- Node.js v16+, https://nodejs.org/
- VS Code v1.71+, https://code.visualstudio.com/
- PNPM v7+, https://pnpm.io/

## Technologies

| Technology | Role                                                 | Sources                        |
| :--------- | :--------------------------------------------------- | :----------------------------- |
| Go         | Backend programming language                         | https://pkg.go.dev/std         |
| TypeScript | Frontend programming language                        | https://typescriptlang.org     |
| Vite       | Next Generation Frontend Tooling                     | https://vitejs.dev/            |
| Vue 3      | Progressive JavaScript Framework                     | https://vuejs.org/             |
| Wails      | Build cross-platform desktop applications using Go   | https://wails.io/              |
| UPX        | Ultimate packer for executables                      | https://upx.github.io/         |
| Swaggo     | Converts Go annotations to Swagger Documentation 2.0 | https://github.com/swaggo/swag |

## Design Pattern Summary

| Pattern   | Role                                                                                     | Packages |
| :-------- | :--------------------------------------------------------------------------------------- | :------- |
| Singleton | Maintains read-only state for the backend included connections to database, logger, etc. | app      |
