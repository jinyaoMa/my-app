import { createApp } from "vue";
import router from "./router";
import i18n from "./i18n";
import Components from "@jinyaoma/my-app-components/index";

import "./style.scss";
import App from "./App.vue";

createApp(App).use(router).use(i18n).use(Components).mount("#app");
