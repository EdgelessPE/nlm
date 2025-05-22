import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Nep from "./views/Nep/index.vue";
import Ept from "./views/Ept.vue";
import Pipeline from "./views/Pipeline/index.vue";
import Storage from "./views/Storage.vue";

const routes = [
  { path: "/", component: Home },
  { path: "/nep", component: Nep },
  {
    path: "/ept",
    component: Ept,
  },
  {
    path: "/pipeline",
    component: Pipeline,
  },
  {
    path: "/storage",
    component: Storage,
  },
];

export default createRouter({
  history: createWebHistory(),
  routes,
});
