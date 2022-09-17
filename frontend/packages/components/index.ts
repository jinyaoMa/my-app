import { App, Plugin } from "vue";
import HelloWorld from "./HelloWorld";

const components = [HelloWorld];

const install = (app: App) => {
  components.map((item) => {
    app.use({
      install: item.install,
    } as Plugin);
  });
};

export default {
  install,
} as Plugin;
