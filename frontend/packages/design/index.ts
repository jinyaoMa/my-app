import type { App, Plugin } from "vue";
import "@jinyaoma/my-app-components/theme-default/style.scss";
import {
  MyIcon,
  MyContainer,
  MyLink,
  MyMenu,
} from "@jinyaoma/my-app-components";

const components = [MyIcon, MyContainer, MyLink, MyMenu];

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
