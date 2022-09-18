import { defineStore } from "pinia";

export const useColorTheme = defineStore("color-theme", {
  state: () => ({
    theme: "system",
  }),
});
