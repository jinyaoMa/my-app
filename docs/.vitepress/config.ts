// import { defineConfigWithTheme } from 'vitepress'
// import { ThemeConfig } from 'your-theme'
import { defineConfig } from "vitepress";

export default defineConfig({
  title: "My Documentation",
  titleTemplate: false,
  description:
    "Everything about and in my application includes introduction, blog posts, technical docs, etc.",
  appearance: true,
  lastUpdated: true,
  markdown: {
    theme: "material-palenight",
    lineNumbers: true,
  },

  locales: {
    "/": {
      lang: "en-US",
      title: "My Documentation",
      description:
        "Everything about and in my application includes introduction, blog posts, technical docs, etc.",
    },
    "/zh/": {
      lang: "zh-CN",
      title: "我的文档",
      description: "所有包含在我的应用中的相关东西，如简介、博客、技术文档等",
    },
  },

  themeConfig: {
    // Type is `DefaultTheme.Config`
    siteTitle: false,
    logo: "/logo.png",
    nav: [
      { text: "指南", link: "/guide/installation.html" },
      { text: "组件", link: "/examples/button.html" },
    ],
    socialLinks: [{ icon: "github", link: "https://gitee.com/geeksdidi" }],
  },

  head: [],

  scrollOffset: 0,

  ignoreDeadLinks: false,

  // set it to subdirectory in production inserting into /backend/web/static/docs/
  base: process.env.NODE_ENV === "production" ? "/docs/" : "/",
  outDir: "../backend/web/docs/",
});
