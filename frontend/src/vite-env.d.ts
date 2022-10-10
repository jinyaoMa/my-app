/// <reference types="vite/client" />

declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

declare module "*.md" {
  import type { ComponentOptions } from "vue";
  const Component: ComponentOptions;
  export default Component;
}

namespace config {
  declare type Config = {
    ColorTheme: string;
    DisplayLanguage: string;
    LogPath: string;
    AutoStart: string;
    PortHttp: string;
    PortHttps: string;
    DirCerts: string;
  };
}
