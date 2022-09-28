import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Test from "@/views/Test.vue";
import Home from "@/views/Home.vue";
import Settings from "@/views/Settings.vue";
import About from "@/views/About.vue";

export const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
  },
  {
    path: "/about",
    name: "About",
    component: About,
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
