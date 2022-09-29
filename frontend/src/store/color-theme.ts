import { defineStore } from "pinia";

const prefix = "my-theme-";

export const useColorTheme = defineStore("color-theme", {
  state: () => ({
    theme: "system",
  }),
  actions: {
    changeTheme(theme: string) {
      const root = document.querySelector(":root");
      if (root) {
        root.classList.remove(prefix + this.theme);
        root.classList.add(prefix + theme);
      }
      this.theme = theme;
    },
  },
});
