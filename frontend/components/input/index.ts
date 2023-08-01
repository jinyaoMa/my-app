import type { App } from "vue";
import Input from "./src/input.vue";

Input.install = (app: App) => {
  app.component(Input.name, Input);
};

export default Input;
export const MyInput = Input;
