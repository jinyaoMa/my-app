declare module "*.vue" {
  import { DefineComponent, Plugin } from "vue";
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any> & Plugin;
  export default component;
}
