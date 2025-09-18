# 快速上手

> ⚠ 建议在初次开发的时候详细阅读该文档

## 安装

```
# install degit
npm  i -g degit

# clone starter
degit windlil/vuecomp-starter [your project name]

# switch to your project directory
cd [your project name]

# install dependence
pnpm i

# open docs
pnpm docs:dev
```

## scripts 命令介绍

- `pnpm build`: 对组件库进行打包。
- `pnpm docs:dev`: 开启说明文档服务。
- `pnpm comp:play`: 开启组件测试环境。
- `pnpm docs:build`: 打包说明文档。

自动化脚本：

- `pnpm comp:create [component name]`: 创建组件文件，并自动引入到导出根文件。
- `pnpm comp:new [component name]`: 创建测试环境组件文件。
- `pnpm docs:new [component name]`: 创建新的组件说明文档。
- `pnpm component:create [component name]`: `pnpm comp:create`、`pnpm comp:new`、`pnpm docs:new`合并命令，推荐使用。

## env 配置介绍

⚠ 在初次开发的时候需要对根目录的.env 文件进行配置修改

```
# CSS预处理器 默认支持sass和less
CSS_PREPROCESSOR = 'scss'
# 用于添加组件命名前缀，建议采用首字母大写格式，最后生成的组件会采取类似这种格式:<my-button></my-button>
COMPONENT_NAME = 'My'
```

## 组件开发约定

组件开发建议采用以下目录结构：

```
├─packages
|    ├─components
|    |     ├─style
|    |     |   └index.scss          // 组件库全局样式
|    |     ├─src
|    |     |  ├─index.ts            // 导出所有组件
|    |     |  ├─button              // 组件文件
|    |     |  |   ├─index.ts        // 在此进行导出和局部注册
|    |     |  |   ├─src             // ❗实际开发中只需要关注到src下的文件
|    |     |  |   |  ├─button.vue
|    |     |  |   |  ├─style
|    |     |  |   |  |   └index.scss
|    |     └─index.ts              // 全局注册所有组件
```

建议采用命令`pnpm component:create [component name]`来创建新组件，帮助自动完成繁琐的引入和创建步骤。

如果不采用自动命令，则按照模板已创建的示例组件的模式来进行开发，这种方法会相当繁琐，因此强烈建议使用命令的方式进行新组件创建！

## 说明文档开发约定

```
├─docs
|  ├─index.md               // 文档首页内容
|  ├─guide
|  |   ├─features.md        // 特性
|  |   ├─started.md         // 快速上手
|  |   ├─components         // 各组件的具体说明
|  |   |     └button.md
```

只需要关注 `docs/guide/components/` 下 md 文件的编写，会自动创建新的侧边栏内容以及路由，不需要关注其它的配置文件。

⚠️ 注意，需要在 md 文件采用以下的格式来确定侧边栏的名称：

```md
---
sider_text="按钮 button"
---
```

组件说明文档具体内容可参考示例组件。

## 自动化部署

具体文件在 `./.github/workflows/deploy.yaml`

```
示例：
git tag docs@v0.0.1

git add .

git commit -m 'chore: update'

git push

git push origin docs@v0.0.1
```
