import { defineStore } from "pinia";

export const useLoading = defineStore("loading", {
  state: () => ({
    loading: true,
  }),
  actions: {
    startLoading() {
      this.loading = true;
    },
    endLoading() {
      this.loading = false;
    },
  },
});
