import { App } from "vue";
import "@jinyaoma/my-app-icons";
import Icon from "./icon.vue";

Icon.install = (app: App) => {
  app.component(Icon.name, Icon);
};

export default Icon;

export const MyIcon = Icon;
