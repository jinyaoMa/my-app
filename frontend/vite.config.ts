import { defineConfig } from "vite";
import { resolve } from "path";
import vue from "@vitejs/plugin-vue";
import vueSetupExtend from "vite-plugin-vue-setup-extend";
import vueI18n from "@intlify/vite-plugin-vue-i18n";
import Markdown from "vite-plugin-md"; // https://github.com/antfu/vite-plugin-md
import MarkdownItPrism from "markdown-it-prism";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({
      include: [/\.vue$/, /\.md$/],
    }),
    vueSetupExtend(),
    vueI18n({
      // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
      compositionOnly: true,
      // you need to set i18n resource including paths !
      include: resolve(__dirname, "src/i18n/locales/**"),
    }),
    Markdown({
      markdownItOptions: {
        html: true,
        linkify: true,
        typographer: true,
      },
      markdownItSetup(md) {
        // add code syntax highlighting with Prism
        md.use(MarkdownItPrism);
      },
    }),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
      "vue-i18n": "vue-i18n/dist/vue-i18n.runtime.esm-bundler.js",
    },
  },
});
