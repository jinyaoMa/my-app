import type { App, Plugin } from 'vue'

type SfcWithInstall<T> = T & Plugin

export const withInstall = <T, E extends Record<string, any>>(main: T, extra?: E) => {
  ;(main as SfcWithInstall<T>).install = (app: App) => {
    for (const comp of [main, ...Object.values(extra ?? {})]) {
      app.component(comp.name, comp)
    }
  }

  if (extra) {
    for (const [key, comp] of Object.entries(extra)) {
      ;(main as any)[key] = comp
    }
  }
  return main as SfcWithInstall<T> & E
}

export const withNoopInstall = <T>(component: T) => {
  ;(component as SfcWithInstall<T>).install = (_app: App) => {}
  return component as SfcWithInstall<T>
}
