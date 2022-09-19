import { App } from "vue";
import Container from "./container.vue";

Container.install = (app: App) => {
  app.component(Container.name, Container);
};

export default Container;

export const MyContainer = Container;
