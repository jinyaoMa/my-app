import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

import Components from "@jinyaoma/my-app-components/index";

createApp(App).use(Components).mount("#app");
