import "@/scss/styles.scss";

import App from "@/App.vue";
import { createApp } from "vue";
import { createPinia } from "pinia";

const app = createApp(App);
const store = createPinia();
app.use(store);
app.mount("#app");
