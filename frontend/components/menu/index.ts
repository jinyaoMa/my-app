import type { App } from "vue";
import Menu from "./src/menu.vue";
import MenuItem from "./src/menu-item.vue";
import MenuGroup from "./src/menu-group.vue";

Menu.install = (app: App) => {
  app.component(Menu.name, Menu);
  app.component(MenuItem.name, MenuItem);
  app.component(MenuGroup.name, MenuGroup);
};

export default Menu;
export const MyMenu = Menu;
export const MyMenuItem = MenuItem;
export const MyMenuGroup = MenuGroup;
