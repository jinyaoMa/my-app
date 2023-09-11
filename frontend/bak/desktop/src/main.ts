import { createApp } from "vue";
import router from "./router";
import store from "./store";
import i18n from "./i18n";
import Components from "@jinyaoma/my-app-design/index";

import "./style.scss";
import App from "./App.vue";

createApp(App).use(router).use(store).use(i18n).use(Components).mount("#app");
