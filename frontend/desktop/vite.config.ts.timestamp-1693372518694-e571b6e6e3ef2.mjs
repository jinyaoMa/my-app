// vite.config.ts
import { defineConfig } from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+vite@4.4.8_@types+node@20.4.5_sass@1.64.2/node_modules/vite/dist/node/index.js";
import { resolve } from "path";
import vue from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+@vitejs+plugin-vue@4.2.3_vite@4.4.8_vue@3.3.4/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import vueSetupExtend from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-vue-setup-extend@0.4.0_vite@4.4.8/node_modules/vite-plugin-vue-setup-extend/dist/index.mjs";
import vueI18n from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+@intlify+unplugin-vue-i18n@0.12.2_vue-i18n@9.2.2/node_modules/@intlify/unplugin-vue-i18n/lib/vite.mjs";
import Markdown from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-md@0.21.5_@vitejs+plugin-vue@4.2.3_sass@1.64.2_vite@4.4.8/node_modules/vite-plugin-md/dist/index.js";
import MarkdownItPrism from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+markdown-it-prism@2.3.0/node_modules/markdown-it-prism/build/index.js";
import MarkdownItLinkAttributes from "file:///C:/Users/jinya/Desktop/my-app/node_modules/.pnpm/registry.npmmirror.com+markdown-it-link-attributes@4.0.1/node_modules/markdown-it-link-attributes/index.js";
var __vite_injected_original_dirname = "C:\\Users\\jinya\\Desktop\\my-app\\frontend\\desktop";
var vite_config_default = defineConfig({
  plugins: [
    vue({
      include: [/\.vue$/, /\.md$/]
    }),
    vueSetupExtend(),
    vueI18n({
      // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
      compositionOnly: true,
      // you need to set i18n resource including paths !
      include: resolve(__vite_injected_original_dirname, "src/i18n/locales/**")
    }),
    Markdown({
      markdownItOptions: {
        html: true,
        linkify: true,
        typographer: true
      },
      markdownItSetup(md) {
        md.use(MarkdownItPrism);
        md.use(MarkdownItLinkAttributes, [
          {
            matcher(href) {
              return href.match(/^(\.|https?:\/\/)/);
            },
            attrs: {
              target: "_blank",
              rel: "noopener"
            }
          }
        ]);
      }
    })
  ],
  resolve: {
    alias: {
      "@": resolve(__vite_injected_original_dirname, "src"),
      "vue-i18n": "vue-i18n/dist/vue-i18n.runtime.esm-bundler.js"
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFxqaW55YVxcXFxEZXNrdG9wXFxcXG15LWFwcFxcXFxmcm9udGVuZFxcXFxkZXNrdG9wXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFxqaW55YVxcXFxEZXNrdG9wXFxcXG15LWFwcFxcXFxmcm9udGVuZFxcXFxkZXNrdG9wXFxcXHZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9DOi9Vc2Vycy9qaW55YS9EZXNrdG9wL215LWFwcC9mcm9udGVuZC9kZXNrdG9wL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSBcInZpdGVcIjtcbmltcG9ydCB7IHJlc29sdmUgfSBmcm9tIFwicGF0aFwiO1xuaW1wb3J0IHZ1ZSBmcm9tIFwiQHZpdGVqcy9wbHVnaW4tdnVlXCI7XG5pbXBvcnQgdnVlU2V0dXBFeHRlbmQgZnJvbSBcInZpdGUtcGx1Z2luLXZ1ZS1zZXR1cC1leHRlbmRcIjtcbmltcG9ydCB2dWVJMThuIGZyb20gXCJAaW50bGlmeS91bnBsdWdpbi12dWUtaTE4bi92aXRlXCI7XG5pbXBvcnQgTWFya2Rvd24gZnJvbSBcInZpdGUtcGx1Z2luLW1kXCI7IC8vIGh0dHBzOi8vZ2l0aHViLmNvbS9hbnRmdS92aXRlLXBsdWdpbi1tZFxuaW1wb3J0IE1hcmtkb3duSXRQcmlzbSBmcm9tIFwibWFya2Rvd24taXQtcHJpc21cIjtcbmltcG9ydCBNYXJrZG93bkl0TGlua0F0dHJpYnV0ZXMgZnJvbSBcIm1hcmtkb3duLWl0LWxpbmstYXR0cmlidXRlc1wiO1xuXG4vLyBodHRwczovL3ZpdGVqcy5kZXYvY29uZmlnL1xuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcbiAgcGx1Z2luczogW1xuICAgIHZ1ZSh7XG4gICAgICBpbmNsdWRlOiBbL1xcLnZ1ZSQvLCAvXFwubWQkL10sXG4gICAgfSksXG4gICAgdnVlU2V0dXBFeHRlbmQoKSxcbiAgICB2dWVJMThuKHtcbiAgICAgIC8vIGlmIHlvdSB3YW50IHRvIHVzZSBWdWUgSTE4biBMZWdhY3kgQVBJLCB5b3UgbmVlZCB0byBzZXQgYGNvbXBvc2l0aW9uT25seTogZmFsc2VgXG4gICAgICBjb21wb3NpdGlvbk9ubHk6IHRydWUsXG4gICAgICAvLyB5b3UgbmVlZCB0byBzZXQgaTE4biByZXNvdXJjZSBpbmNsdWRpbmcgcGF0aHMgIVxuICAgICAgaW5jbHVkZTogcmVzb2x2ZShfX2Rpcm5hbWUsIFwic3JjL2kxOG4vbG9jYWxlcy8qKlwiKSxcbiAgICB9KSxcbiAgICBNYXJrZG93bih7XG4gICAgICBtYXJrZG93bkl0T3B0aW9uczoge1xuICAgICAgICBodG1sOiB0cnVlLFxuICAgICAgICBsaW5raWZ5OiB0cnVlLFxuICAgICAgICB0eXBvZ3JhcGhlcjogdHJ1ZSxcbiAgICAgIH0sXG4gICAgICBtYXJrZG93bkl0U2V0dXAobWQpIHtcbiAgICAgICAgLy8gYWRkIGNvZGUgc3ludGF4IGhpZ2hsaWdodGluZyB3aXRoIFByaXNtXG4gICAgICAgIG1kLnVzZShNYXJrZG93bkl0UHJpc20pO1xuICAgICAgICBtZC51c2UoTWFya2Rvd25JdExpbmtBdHRyaWJ1dGVzLCBbXG4gICAgICAgICAge1xuICAgICAgICAgICAgbWF0Y2hlcihocmVmKSB7XG4gICAgICAgICAgICAgIHJldHVybiBocmVmLm1hdGNoKC9eKFxcLnxodHRwcz86XFwvXFwvKS8pO1xuICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIGF0dHJzOiB7XG4gICAgICAgICAgICAgIHRhcmdldDogXCJfYmxhbmtcIixcbiAgICAgICAgICAgICAgcmVsOiBcIm5vb3BlbmVyXCIsXG4gICAgICAgICAgICB9LFxuICAgICAgICAgIH0sXG4gICAgICAgIF0pO1xuICAgICAgfSxcbiAgICB9KSxcbiAgXSxcbiAgcmVzb2x2ZToge1xuICAgIGFsaWFzOiB7XG4gICAgICBcIkBcIjogcmVzb2x2ZShfX2Rpcm5hbWUsIFwic3JjXCIpLFxuICAgICAgXCJ2dWUtaTE4blwiOiBcInZ1ZS1pMThuL2Rpc3QvdnVlLWkxOG4ucnVudGltZS5lc20tYnVuZGxlci5qc1wiLFxuICAgIH0sXG4gIH0sXG59KTtcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBMFUsU0FBUyxvQkFBb0I7QUFDdlcsU0FBUyxlQUFlO0FBQ3hCLE9BQU8sU0FBUztBQUNoQixPQUFPLG9CQUFvQjtBQUMzQixPQUFPLGFBQWE7QUFDcEIsT0FBTyxjQUFjO0FBQ3JCLE9BQU8scUJBQXFCO0FBQzVCLE9BQU8sOEJBQThCO0FBUHJDLElBQU0sbUNBQW1DO0FBVXpDLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLFNBQVM7QUFBQSxJQUNQLElBQUk7QUFBQSxNQUNGLFNBQVMsQ0FBQyxVQUFVLE9BQU87QUFBQSxJQUM3QixDQUFDO0FBQUEsSUFDRCxlQUFlO0FBQUEsSUFDZixRQUFRO0FBQUE7QUFBQSxNQUVOLGlCQUFpQjtBQUFBO0FBQUEsTUFFakIsU0FBUyxRQUFRLGtDQUFXLHFCQUFxQjtBQUFBLElBQ25ELENBQUM7QUFBQSxJQUNELFNBQVM7QUFBQSxNQUNQLG1CQUFtQjtBQUFBLFFBQ2pCLE1BQU07QUFBQSxRQUNOLFNBQVM7QUFBQSxRQUNULGFBQWE7QUFBQSxNQUNmO0FBQUEsTUFDQSxnQkFBZ0IsSUFBSTtBQUVsQixXQUFHLElBQUksZUFBZTtBQUN0QixXQUFHLElBQUksMEJBQTBCO0FBQUEsVUFDL0I7QUFBQSxZQUNFLFFBQVEsTUFBTTtBQUNaLHFCQUFPLEtBQUssTUFBTSxtQkFBbUI7QUFBQSxZQUN2QztBQUFBLFlBQ0EsT0FBTztBQUFBLGNBQ0wsUUFBUTtBQUFBLGNBQ1IsS0FBSztBQUFBLFlBQ1A7QUFBQSxVQUNGO0FBQUEsUUFDRixDQUFDO0FBQUEsTUFDSDtBQUFBLElBQ0YsQ0FBQztBQUFBLEVBQ0g7QUFBQSxFQUNBLFNBQVM7QUFBQSxJQUNQLE9BQU87QUFBQSxNQUNMLEtBQUssUUFBUSxrQ0FBVyxLQUFLO0FBQUEsTUFDN0IsWUFBWTtBQUFBLElBQ2Q7QUFBQSxFQUNGO0FBQ0YsQ0FBQzsiLAogICJuYW1lcyI6IFtdCn0K
