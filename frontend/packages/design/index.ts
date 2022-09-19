import { App, Plugin } from "vue";
import "@jinyaoma/my-app-components/theme/default.scss";
import { MyHelloWorld, MyIcon } from "@jinyaoma/my-app-components";

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
