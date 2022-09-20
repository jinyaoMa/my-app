import type { App } from "vue";
import Container from "./src/container.vue";
import Aside from "./src/aside.vue";
import Main from "./src/main.vue";
import Header from "./src/header.vue";
import Footer from "./src/footer.vue";

Container.install = (app: App) => {
  app.component(Container.name, Container);
  app.component(Aside.name, Aside);
  app.component(Main.name, Main);
  app.component(Header.name, Header);
  app.component(Footer.name, Footer);
};
Container.Aside = Aside;
Container.Main = Main;
Container.Header = Header;
Container.Footer = Footer;

export default Container;
export const MyContainer = Container;
export const MyAside = Aside;
export const MyMain = Main;
export const MyHeader = Header;
export const MyFooter = Footer;
