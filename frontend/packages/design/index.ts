import { App, Plugin } from "vue";
import { MyHelloWorld, MyIcon } from "../components";

const components = [MyHelloWorld, MyIcon];

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
