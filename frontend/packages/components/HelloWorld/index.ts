import { App } from "vue";
import HelloWorld from "./HelloWorld.vue";

HelloWorld.install = (app: App) => {
  app.component(HelloWorld.name, HelloWorld);
};

export const MyHelloWorld = HelloWorld;

export default HelloWorld;
