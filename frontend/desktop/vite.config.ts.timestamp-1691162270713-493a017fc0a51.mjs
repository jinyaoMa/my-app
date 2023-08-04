// vite.config.ts
import { defineConfig } from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+vite@4.4.8_@types+node@20.4.5_sass@1.64.2/node_modules/vite/dist/node/index.js";
import { resolve } from "path";
import vue from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+@vitejs+plugin-vue@4.2.3_vite@4.4.8_vue@3.3.4/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import vueSetupExtend from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-vue-setup-extend@0.4.0_vite@4.4.8/node_modules/vite-plugin-vue-setup-extend/dist/index.mjs";
import vueI18n from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+@intlify+unplugin-vue-i18n@0.12.2_vue-i18n@9.2.2/node_modules/@intlify/unplugin-vue-i18n/lib/vite.mjs";
import Markdown from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-md@0.21.5_@vitejs+plugin-vue@4.2.3_sass@1.64.2_vite@4.4.8/node_modules/vite-plugin-md/dist/index.js";
import MarkdownItPrism from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+markdown-it-prism@2.3.0/node_modules/markdown-it-prism/build/index.js";
import MarkdownItLinkAttributes from "file:///E:/GitHub/my-app/node_modules/.pnpm/registry.npmmirror.com+markdown-it-link-attributes@4.0.1/node_modules/markdown-it-link-attributes/index.js";
var __vite_injected_original_dirname = "E:\\GitHub\\my-app\\frontend\\desktop";
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
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJFOlxcXFxHaXRIdWJcXFxcbXktYXBwXFxcXGZyb250ZW5kXFxcXGRlc2t0b3BcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIkU6XFxcXEdpdEh1YlxcXFxteS1hcHBcXFxcZnJvbnRlbmRcXFxcZGVza3RvcFxcXFx2aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovR2l0SHViL215LWFwcC9mcm9udGVuZC9kZXNrdG9wL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSBcInZpdGVcIjtcclxuaW1wb3J0IHsgcmVzb2x2ZSB9IGZyb20gXCJwYXRoXCI7XHJcbmltcG9ydCB2dWUgZnJvbSBcIkB2aXRlanMvcGx1Z2luLXZ1ZVwiO1xyXG5pbXBvcnQgdnVlU2V0dXBFeHRlbmQgZnJvbSBcInZpdGUtcGx1Z2luLXZ1ZS1zZXR1cC1leHRlbmRcIjtcclxuaW1wb3J0IHZ1ZUkxOG4gZnJvbSBcIkBpbnRsaWZ5L3VucGx1Z2luLXZ1ZS1pMThuL3ZpdGVcIjtcclxuaW1wb3J0IE1hcmtkb3duIGZyb20gXCJ2aXRlLXBsdWdpbi1tZFwiOyAvLyBodHRwczovL2dpdGh1Yi5jb20vYW50ZnUvdml0ZS1wbHVnaW4tbWRcclxuaW1wb3J0IE1hcmtkb3duSXRQcmlzbSBmcm9tIFwibWFya2Rvd24taXQtcHJpc21cIjtcclxuaW1wb3J0IE1hcmtkb3duSXRMaW5rQXR0cmlidXRlcyBmcm9tIFwibWFya2Rvd24taXQtbGluay1hdHRyaWJ1dGVzXCI7XHJcblxyXG4vLyBodHRwczovL3ZpdGVqcy5kZXYvY29uZmlnL1xyXG5leHBvcnQgZGVmYXVsdCBkZWZpbmVDb25maWcoe1xyXG4gIHBsdWdpbnM6IFtcclxuICAgIHZ1ZSh7XHJcbiAgICAgIGluY2x1ZGU6IFsvXFwudnVlJC8sIC9cXC5tZCQvXSxcclxuICAgIH0pLFxyXG4gICAgdnVlU2V0dXBFeHRlbmQoKSxcclxuICAgIHZ1ZUkxOG4oe1xyXG4gICAgICAvLyBpZiB5b3Ugd2FudCB0byB1c2UgVnVlIEkxOG4gTGVnYWN5IEFQSSwgeW91IG5lZWQgdG8gc2V0IGBjb21wb3NpdGlvbk9ubHk6IGZhbHNlYFxyXG4gICAgICBjb21wb3NpdGlvbk9ubHk6IHRydWUsXHJcbiAgICAgIC8vIHlvdSBuZWVkIHRvIHNldCBpMThuIHJlc291cmNlIGluY2x1ZGluZyBwYXRocyAhXHJcbiAgICAgIGluY2x1ZGU6IHJlc29sdmUoX19kaXJuYW1lLCBcInNyYy9pMThuL2xvY2FsZXMvKipcIiksXHJcbiAgICB9KSxcclxuICAgIE1hcmtkb3duKHtcclxuICAgICAgbWFya2Rvd25JdE9wdGlvbnM6IHtcclxuICAgICAgICBodG1sOiB0cnVlLFxyXG4gICAgICAgIGxpbmtpZnk6IHRydWUsXHJcbiAgICAgICAgdHlwb2dyYXBoZXI6IHRydWUsXHJcbiAgICAgIH0sXHJcbiAgICAgIG1hcmtkb3duSXRTZXR1cChtZCkge1xyXG4gICAgICAgIC8vIGFkZCBjb2RlIHN5bnRheCBoaWdobGlnaHRpbmcgd2l0aCBQcmlzbVxyXG4gICAgICAgIG1kLnVzZShNYXJrZG93bkl0UHJpc20pO1xyXG4gICAgICAgIG1kLnVzZShNYXJrZG93bkl0TGlua0F0dHJpYnV0ZXMsIFtcclxuICAgICAgICAgIHtcclxuICAgICAgICAgICAgbWF0Y2hlcihocmVmKSB7XHJcbiAgICAgICAgICAgICAgcmV0dXJuIGhyZWYubWF0Y2goL14oXFwufGh0dHBzPzpcXC9cXC8pLyk7XHJcbiAgICAgICAgICAgIH0sXHJcbiAgICAgICAgICAgIGF0dHJzOiB7XHJcbiAgICAgICAgICAgICAgdGFyZ2V0OiBcIl9ibGFua1wiLFxyXG4gICAgICAgICAgICAgIHJlbDogXCJub29wZW5lclwiLFxyXG4gICAgICAgICAgICB9LFxyXG4gICAgICAgICAgfSxcclxuICAgICAgICBdKTtcclxuICAgICAgfSxcclxuICAgIH0pLFxyXG4gIF0sXHJcbiAgcmVzb2x2ZToge1xyXG4gICAgYWxpYXM6IHtcclxuICAgICAgXCJAXCI6IHJlc29sdmUoX19kaXJuYW1lLCBcInNyY1wiKSxcclxuICAgICAgXCJ2dWUtaTE4blwiOiBcInZ1ZS1pMThuL2Rpc3QvdnVlLWkxOG4ucnVudGltZS5lc20tYnVuZGxlci5qc1wiLFxyXG4gICAgfSxcclxuICB9LFxyXG59KTtcclxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUErUixTQUFTLG9CQUFvQjtBQUM1VCxTQUFTLGVBQWU7QUFDeEIsT0FBTyxTQUFTO0FBQ2hCLE9BQU8sb0JBQW9CO0FBQzNCLE9BQU8sYUFBYTtBQUNwQixPQUFPLGNBQWM7QUFDckIsT0FBTyxxQkFBcUI7QUFDNUIsT0FBTyw4QkFBOEI7QUFQckMsSUFBTSxtQ0FBbUM7QUFVekMsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDMUIsU0FBUztBQUFBLElBQ1AsSUFBSTtBQUFBLE1BQ0YsU0FBUyxDQUFDLFVBQVUsT0FBTztBQUFBLElBQzdCLENBQUM7QUFBQSxJQUNELGVBQWU7QUFBQSxJQUNmLFFBQVE7QUFBQTtBQUFBLE1BRU4saUJBQWlCO0FBQUE7QUFBQSxNQUVqQixTQUFTLFFBQVEsa0NBQVcscUJBQXFCO0FBQUEsSUFDbkQsQ0FBQztBQUFBLElBQ0QsU0FBUztBQUFBLE1BQ1AsbUJBQW1CO0FBQUEsUUFDakIsTUFBTTtBQUFBLFFBQ04sU0FBUztBQUFBLFFBQ1QsYUFBYTtBQUFBLE1BQ2Y7QUFBQSxNQUNBLGdCQUFnQixJQUFJO0FBRWxCLFdBQUcsSUFBSSxlQUFlO0FBQ3RCLFdBQUcsSUFBSSwwQkFBMEI7QUFBQSxVQUMvQjtBQUFBLFlBQ0UsUUFBUSxNQUFNO0FBQ1oscUJBQU8sS0FBSyxNQUFNLG1CQUFtQjtBQUFBLFlBQ3ZDO0FBQUEsWUFDQSxPQUFPO0FBQUEsY0FDTCxRQUFRO0FBQUEsY0FDUixLQUFLO0FBQUEsWUFDUDtBQUFBLFVBQ0Y7QUFBQSxRQUNGLENBQUM7QUFBQSxNQUNIO0FBQUEsSUFDRixDQUFDO0FBQUEsRUFDSDtBQUFBLEVBQ0EsU0FBUztBQUFBLElBQ1AsT0FBTztBQUFBLE1BQ0wsS0FBSyxRQUFRLGtDQUFXLEtBQUs7QUFBQSxNQUM3QixZQUFZO0FBQUEsSUFDZDtBQUFBLEVBQ0Y7QUFDRixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
