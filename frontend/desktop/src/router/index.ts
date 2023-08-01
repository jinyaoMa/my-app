import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import Settings from "@/views/Settings.vue";
import About from "@/views/About.vue";
import Test from "@/views/Test.vue";
import Keyring from "@/views/Keyring.vue";

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
  {
    path: "/keyring",
    name: "Keyring",
    component: Keyring,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
