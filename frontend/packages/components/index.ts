import { App } from "vue";
import HelloWorld from "./HelloWorld.vue";

const components = [HelloWorld];

export default {
  install(app: App) {
    components.forEach((component) => {
      app.use(component);
    });
  },
  ...components,
};
