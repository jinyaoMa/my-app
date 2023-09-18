import type { App } from "vue";
import type { SFCWithInstall } from "./types";

export const withInstall = <T extends Record<string, any>>(main: T) => {
  (main as SFCWithInstall<T>).install = (app: App): void => {
    app.component(main.name, main);
  };

  return main as SFCWithInstall<T>;
};
