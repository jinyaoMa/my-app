import { App } from "vue";
import Icon from "./icon.vue";

Icon.install = (app: App) => {
  app.component(Icon.name, Icon);
};

export default Icon;

export const MyIcon = Icon;
