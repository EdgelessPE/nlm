import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Nep from "./views/Nep.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/nep", component: Nep },
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
