import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
// import Home from "@/views/Home.vue";

export const routes: Array<RouteRecordRaw> = [
  /*
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      icon: "home",
    },
  },
  */
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
