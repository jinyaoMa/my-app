import type { App } from "vue";
import Button from "./src/button.vue";

Button.install = (app: App) => {
  app.component(Button.name, Button);
};

export default Button;
export const MyButton = Button;
