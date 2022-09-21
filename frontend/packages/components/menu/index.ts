import type { App } from "vue";
import Menu from "./src/menu.vue";

Menu.install = (app: App) => {
  app.component(Menu.name, Menu);
};

export default Menu;
export const MyMenu = Menu;
