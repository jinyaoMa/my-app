import type { App } from "vue";
import Select from "./src/select.vue";
import Option from "./src/option.vue";

Select.install = (app: App) => {
  app.component(Select.name, Select);
  app.component(Option.name, Option);
};

export default Select;
export const MySelect = Select;
export const MyOption = Option;
