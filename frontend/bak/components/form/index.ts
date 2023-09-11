import type { App } from "vue";
import Form from "./src/form.vue";
import FormItem from "./src/form-item.vue";
import FormGroup from "./src/form-group.vue";

Form.install = (app: App) => {
  app.component(Form.name, Form);
  app.component(FormItem.name, FormItem);
  app.component(FormGroup.name, FormGroup);
};

export default Form;
export const MyForm = Form;
export const MyFormItem = FormItem;
export const MyFormGroup = FormGroup;
