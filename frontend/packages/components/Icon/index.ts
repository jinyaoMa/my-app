import { App } from "vue";
import Icon from "./Icon.vue";

Icon.install = (app: App) => {
  app.component(Icon.name, Icon);
};

export const MyIcon = Icon;

export default Icon;
