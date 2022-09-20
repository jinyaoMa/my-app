import type { App } from "vue";
import Container from "./src/container.vue";
import Aside from "./src/aside.vue";
import Main from "./src/main.vue";

Container.install = (app: App) => {
  app.component(Container.name, Container);
  app.component(Aside.name, Aside);
  app.component(Main.name, Main);
};
Container.Aside = Aside;
Container.Main = Main;

export default Container;
export const MyContainer = Container;
export const MyAside = Aside;
export const MyMain = Main;
