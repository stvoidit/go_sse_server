import { createApp } from "vue";
import App from "@/App.vue";
import { createPinia } from "pinia";
import "@/scss/styles.scss";
const app = createApp(App);
const store = createPinia();
app.use(store);
app.mount("#app");
