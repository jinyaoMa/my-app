import { App } from "vue";
import HelloWorld from "./hello-world.vue";

HelloWorld.install = (app: App) => {
  app.component(HelloWorld.name, HelloWorld);
};

export default HelloWorld;

export const MyHelloWorld = HelloWorld;
