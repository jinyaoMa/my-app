import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Test from "@/views/Test.vue";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Test,
  },
  {
    path: "/settings",
    name: "Settings",
    component: Test,
  },
  {
    path: "/about",
    name: "About",
    component: Test,
  },
  {
    path: "/test",
    name: "Test",
    component: Test,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
